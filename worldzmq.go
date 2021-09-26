package main

import (
	"log"
	"strings"

	types "github.com/gfx-labs/etherlands/types"
	"github.com/zeromq/goczmq"
)

type WorldZmq struct {
	W *types.World

	publisher  *goczmq.Channeler
	subscriber *goczmq.Channeler
}

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
		message := <-Z.subscriber.RecvChan
		command := string(message[0])
		args := strings.Split(string(message[1]), ":")
		switch command {
		case "GET":
			log.Println(args)
		default:
			log.Println("Unrecognized Command:", command)
		}
	}
}
