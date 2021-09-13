package main

import (
	"context"
	"fmt"

	"github.com/gfx-labs/etherlands/types"
	"github.com/mediocregopher/radix/v4"
)


type MemoryCache struct {
	redis radix.Client
	ctx *context.Context
}

func NewMemoryCache(ctx *context.Context) (*MemoryCache, error) {

	redis, err := (radix.PoolConfig{}).New(*ctx,"tcp","127.0.0.1:6379")
	if(err != nil){
		return nil, err
	}

	return &MemoryCache{redis:redis, ctx:ctx}, nil

}

func (M *MemoryCache) CachePlot(plot *types.Plot){
	key_x := fmt.Sprintf("plot:%d:x",plot.ChainId())
	value_x := fmt.Sprintf("%d",plot.X())
	key_z := fmt.Sprintf("plot:%d:z",plot.ChainId())
	value_z := fmt.Sprintf("%d",plot.Z())
	key_coord := fmt.Sprintf("plot_coord:%d_%d",plot.X(),plot.Z())
	value_coord := fmt.Sprintf("%d",plot.ChainId())
	M.redis.Do(*M.ctx, radix.Cmd(nil,"MSET",key_x, value_x, key_z, value_z,key_coord,value_coord))
}
