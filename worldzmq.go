package main

import (
	"log"

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
	subscriber := goczmq.NewSubChanneler("tcp://127.0.0.1:10106")
	zmq := WorldZmq{W: world, publisher: publisher, subscriber: subscriber}
	go zmq.StartPublishing()
	go zmq.StartListening()
	return nil
}

func (Z *WorldZmq) StartPublishing() {
}

func (Z *WorldZmq) StartListening() {
	for {
		message := <-Z.subscriber.RecvChan
		log.Println(message)
	}
}
