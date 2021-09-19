package main

import (
	"context"

	"github.com/mediocregopher/radix/v4"
)


type Broker struct {
	redis radix.Client
	ctx *context.Context
}

func NewBroker(ctx *context.Context) (*Broker, error) {

	redis, err := (radix.PoolConfig{}).New(*ctx,"tcp","127.0.0.1:6379")
	if(err != nil){
		return nil, err
	}

	return &Broker{redis:redis, ctx:ctx}, nil

}

func (B *Broker) PublishLink(body string) (error){
	return B.redis.Do(*B.ctx,radix.Cmd(nil,"PUBLISH",body))
}
