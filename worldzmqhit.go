package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func (Z *WorldZmq) hit_scope(args VarArgs) {
	scope, err := args.MustGet(0)
	if Z.checkError(args, err) {
		return
	}

	switch scope {
	case "world":
		Z.hit_world_type(args)
	default:
		Z.checkError(args, errors.New("Unspecified Scope: "+scope))
	}
}

func (Z *WorldZmq) world_link_request(args VarArgs) error {
	uuid_str, err := args.MustGet(2)
	if err != nil {
		return err
	}
	_, err = uuid.Parse(uuid_str)
	if err != nil {
		return errors.New(fmt.Sprintf("malformed uuid %s", uuid_str))
	}
	a, err := args.MustGet(3)
	if err != nil {
		return err
	}
	b, err := args.MustGet(4)
	if err != nil {
		return err
	}
	c, err := args.MustGet(5)
	if err != nil {
		return err
	}
	if a != "" && b != "" && c != "" {
		Z.W.CreateLinkRequest(fmt.Sprintf("%s:%s:%s:%s", uuid_str, a, b, c))
		return nil
	}
	return errors.New(fmt.Sprintf("invalid input %v", args))
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

func (Z *WorldZmq) opensea_image_download(contract, id string) error {
	unlock := Z.lock("nft" + contract + ":" + id)
	defer unlock()
	var url string
	if strings.HasPrefix(contract, "0x") {
		url = fmt.Sprintf(
			"https://api.opensea.io/api/v1/assets?asset_contract_address=%s&token_ids=%s",
			contract,
			id,
		)
	} else {
		url = fmt.Sprintf("https://api.opensea.io/api/v1/assets?collection=%s&token_ids=%s", contract, id)
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
		file, err := os.Create(path.Join(folder, contract))
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, resp_img.Body)
		return err
	}
	return errors.New("no image found")
}

func (Z *WorldZmq) hit_world_type(args VarArgs) {
	dtype, err := args.MustGet(1)
	if Z.checkError(args, err) {
		return
	}
	switch dtype {
	case "gamer":
		Z.hit_world_gamer_field(args)
	case "plot":
		Z.hit_world_plot_field(args)
	case "district":
		Z.hit_world_district_field(args)
	case "town":
		Z.hit_world_town_field(args)
	case "link_request":
		err = Z.world_link_request(args)
		if Z.checkError(args, err) {
			return
		}
	case "image_download":
		collection, err := args.MustGet(2)
		if Z.checkError(args, err) {
			return
		}
		nft_id, err := args.MustGet(3)
		if Z.checkError(args, err) {
			return
		}
		err = Z.opensea_image_download(collection, nft_id)
		if Z.checkError(args, err) {
			return
		}
	default:
		Z.checkError(args, errors.New("Unspecified Type: "+dtype))
	}
}

func (Z *WorldZmq) hit_world_plot_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	uuid_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	plot_id, err := strconv.ParseUint(uuid_str, 10, 64)
	if Z.checkError(args, err) {
		return
	}
	_, err = Z.W.GetPlot(plot_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	default:
		Z.genericError(args, field)
	}
}
func (Z *WorldZmq) hit_world_gamer_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	uuid_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	gamer_id, err := uuid.Parse(uuid_str)
	if Z.checkError(args, err) {
		return
	}
	gamer := Z.W.GetGamer(gamer_id)
	switch field {
	case "pos":
		x, err := args.MustGetInt64(4)
		if Z.checkError(args, err) {
			return
		}
		y, err := args.MustGetInt64(5)
		if Z.checkError(args, err) {
			return
		}
		z, err := args.MustGetInt64(6)
		if Z.checkError(args, err) {
			return
		}
		gamer.SetPosXYZ(x, y, z)
		Z.sendResponse(args, "true")
	case "create_town":
		name, err := args.MustGet(4)
		if Z.checkError(args, err) {
			return
		}
		if !gamer.HasTown() {
			go Z.W.CreateTown(name, gamer)
			Z.sendResponse(args, "true")
		} else {
			Z.sendResponse(args, "false")
		}
	default:
		Z.genericError(args, field)
	}
}

func (Z *WorldZmq) hit_world_district_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	district_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	district_id, err := strconv.ParseUint(district_str, 10, 64)
	if Z.checkError(args, err) {
		return
	}
	district, err := Z.W.GetDistrict(district_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "name":
		Z.sendResponse(args, district.StringName())
	case "plots":
		plots := Z.W.PlotsOfDistrict(district.DistrictId())
		temp := make([]string, len(plots))
		for i := 0; i < len(plots); i++ {
			temp[i] = strconv.FormatUint(plots[i], 10)
		}
		Z.sendResponse(args, strings.Join(temp, "_"))
	case "clusters":
		clusters := Z.W.Cache().GetClusters(district.DistrictId())
		value := ""
		for _, cluster := range clusters {
			value = value + fmt.Sprintf(
				"%d:%d:%d",
				cluster.OriginX,
				cluster.OriginZ,
				len(cluster.Offsets),
			)
			value = value + "@"
		}
		if len(value) > 1 {
			value = value[:len(value)-1]
		}
		Z.sendResponse(args, value)
	case "owner_addr":
		Z.sendResponse(args, district.OwnerAddress())
	default:
		Z.genericError(args, field)
	}
}

func (Z *WorldZmq) hit_world_town_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	town_name, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	town, err := Z.W.GetTown(town_name)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "name":
		Z.sendResponse(args, town.Name())
	case "owner_uuid":
		Z.sendResponse(args, town.Owner().String())
	default:
		Z.genericError(args, field)
	}
}
