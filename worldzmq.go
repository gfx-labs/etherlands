package main

import (
	"errors"
	"log"
	"strings"

	types "github.com/gfx-labs/etherlands/types"
	"github.com/google/uuid"
	"github.com/zeromq/goczmq"
)

type WorldZmq struct {
	W *types.World

	publisher  *goczmq.Channeler
	subscriber *goczmq.Channeler
}

type VarArgs []string

func StartWorldZmq(world *types.World) error {
	publisher := goczmq.NewPubChanneler("tcp://*:10105")
	subscriber := goczmq.NewSubChanneler("tcp://127.0.0.1:10106", "")
	zmq := &WorldZmq{W: world, publisher: publisher, subscriber: subscriber}
	go zmq.StartPublishing()
	go zmq.StartListening()
	return nil
}

func (Z *WorldZmq) StartPublishing() {
}

func (Z *WorldZmq) StartListening() {
	log.Println("now listening")
	for {
		var args VarArgs
		message := <-Z.subscriber.RecvChan
		verb := string(message[0])
		args = strings.Split(string(message[1]), ":")
		log.Printf("[%s] %s\n", verb, args.Command())
		switch verb {
		case "GET":
			Z.get_scope(args)
		default:
			log.Println("Unrecognized Verb:", verb)
		}
	}
}

func (Z *WorldZmq) get_scope(args VarArgs) {
	scope, err := args.MustGet(0)
	if Z.checkError(args.Command(), err) {
		return
	}

	switch scope {
	case "world":
		Z.get_world_type(args)
	default:
		Z.checkError(args.Command(), errors.New("Unspecified Scope: "+scope))
	}
}

func (Z *WorldZmq) get_world_type(args VarArgs) {
	dtype, err := args.MustGet(1)
	if Z.checkError(args.Command(), err) {
		return
	}
	switch dtype {
	case "gamer":
		Z.get_world_gamer_field(args)
	case "links":
		addr, err := args.MustGet(2)
		if Z.checkError(args.Command(), err) {
			return
		}
		gamer_str, err := Z.W.Cache().GetLink(addr)
		if Z.checkError(args.Command(), err) {
			return
		}
		Z.sendResponse(args.Command(), gamer_str)
	default:
		Z.checkError(args.Command(), errors.New("Unspecified Type: "+dtype))
	}
}

func (Z *WorldZmq) get_world_gamer_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args.Command(), err) {
		return
	}
	uuid_str, err := args.MustGet(2)
	if Z.checkError(args.Command(), err) {
		return
	}
	gamer_id, err := uuid.Parse(uuid_str)
	if Z.checkError(args.Command(), err) {
		return
	}
	gamer := Z.W.GetGamer(gamer_id)
	switch field {
	case "address":
		if gamer.Address() != "" {
			Z.sendResponse(args.Command(), gamer.Address())
		}
	default:
		Z.checkError(args.Command(), errors.New("Unspecified Field: "+field))
	}
}

func (Z *WorldZmq) sendResponse(key string, content string) {
	Z.publisher.SendChan <- [][]byte{
		[]byte(key),
		[]byte(content),
	}
}

func (Z *WorldZmq) checkError(key string, err error) bool {
	if err != nil {
		Z.publisher.SendChan <- [][]byte{
			[]byte(key),
			[]byte("error:" + err.Error()),
		}
		return true
	}
	return false
}

func (Args *VarArgs) Command() string {
	return strings.Join(*Args, ":")
}

func (Args *VarArgs) MustGet(idx int) (string, error) {
	if len(*Args) > idx {
		return (*Args)[idx], nil
	}
	return "", errors.New("Variable out of bounds")
}

func (Args *VarArgs) MightGet(idx int) (string, bool) {
	if len(*Args) > idx {
		return (*Args)[idx], true
	}
	return "", false
}
