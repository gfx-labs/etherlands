package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"

	types "github.com/gfx-labs/etherlands/types"
	"github.com/google/uuid"
	"github.com/mediocregopher/radix/v4"
)

type WorldSMP struct {
	W     *types.World
	redis radix.PubSubConn
	ctx   *context.Context

	mutexes sync.Map
}

func StartWorldSMP(world *types.World) error {
	ctx := context.Background()
	config := radix.PersistentPubSubConfig{}
	redispubsub, err := config.New(ctx, func() (network string, addr string, err error) {
		return "tcp", "127.0.0.1:6379", nil
	})
	if err != nil {
		return err
	}
	smp := WorldSMP{W: world, redis: redispubsub, ctx: &ctx}
	err = smp.start()
	return err
}

func (C *WorldSMP) lock(name string) func() {
	value, _ := C.mutexes.LoadOrStore(name, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

func (C *WorldSMP) start() error {
	a := make(chan radix.PubSubMessage, 20)
	C.redis.PSubscribe(*C.ctx, a, "smp:*")
	go func() {
		for {
			payload := <-a
			log.Println("[smp]", payload.Channel, string(payload.Message))
			switch payload.Channel {
			case "smp:image_download":
				// message format: 0xaddr:00id
				go func() {
					err := C.smp_image_download(string(payload.Message))
					if err != nil {
						log.Println("smp:image_download", err)
					}
				}()
			case "smp:link_request":
				// message format: uuid-abc:worda:wordb:wordc
				go func() {
					err := C.smp_link_request(string(payload.Message))
					if err != nil {
						log.Println("smp:link_request", err)
					}
				}()
			default:
				log.Printf("No handler for %s\n", payload.Channel)
			}
		}
	}()
	return errors.New("SMP Event Loop ended")
}

func (C *WorldSMP) smp_link_request(message string) error {
	args := strings.Split(message, ":")
	if len(args) == 4 {
		_, err := uuid.Parse(args[0])
		if err != nil {
			return errors.New(fmt.Sprintf("malformed uuid %s", args[0]))
		}
		if args[1] != "" && args[2] != "" && args[3] != "" {
			C.W.CreateLinkRequest(message)
		}
	}
	return errors.New(fmt.Sprintf("invalid input %s %v", message, args))
}

type opensea_asset_response struct {
	Assets []struct {
		ImageURL      string `json:"image_url"`
		AssetContract struct {
			Address string `json:"address"`
		} `json:"asset_contract"`
		Owner struct {
			Address string `json:"address"`
		} `json:"owner"`
		Collection struct {
			Slug string `json:"slug"`
		} `json:"collection"`
	} `json:"assets"`
}

func (C *WorldSMP) smp_image_download(message string) error {
	args := strings.Split(message, ":")
	if len(args) == 2 {
		unlock := C.lock(message)
		defer unlock()
		var url string
		if strings.HasPrefix(args[0], "0x") {
			url = fmt.Sprintf(
				"https://api.opensea.io/api/v1/assets?asset_contract_address=%s&token_ids=%s",
				args[0],
				args[1],
			)
		} else {
			url = fmt.Sprintf("https://api.opensea.io/api/v1/assets?collection=%s&token_ids=%s", args[0], args[1])
		}
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return err
		}
		resp_bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var formatted opensea_asset_response
		err = json.Unmarshal(resp_bytes, &formatted)
		if err != nil {
			return err
		}
		log.Println(formatted)
		if len(formatted.Assets) > 0 {
			image_url := formatted.Assets[0].ImageURL
			resp_img, err := http.Get(image_url)
			if err != nil {
				return err
			}
			defer resp_img.Body.Close()
			folder := path.Join(
				"./db",
				"images",
				"opensea",
				formatted.Assets[0].AssetContract.Address,
			)
			symlink := path.Join("./db", "images", "opensea", formatted.Assets[0].Collection.Slug)
			os.MkdirAll(folder, 0777)
			if _, err = os.Stat(symlink); os.IsNotExist(err) {
				os.Symlink(formatted.Assets[0].AssetContract.Address, symlink)
			}
			file, err := os.Create(path.Join(folder, args[1]))
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, resp_img.Body)
			return err
		}
		return errors.New(fmt.Sprintf("no image found for %s", message))
	}
	return errors.New(fmt.Sprintf("invalid input %s", message))
}
