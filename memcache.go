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
	key_x := fmt.Sprintf("plot:%d:x",plot.PlotId())
	value_x := fmt.Sprintf("%d",plot.X())
	key_z := fmt.Sprintf("plot:%d:z",plot.PlotId())
	value_z := fmt.Sprintf("%d",plot.Z())
	key_coord := fmt.Sprintf("plot_coord:%d_%d",plot.X(),plot.Z())
	value_coord := fmt.Sprintf("%d",plot.PlotId())
	//key_district := fmt.Sprintf("plot:%d:district",plot.PlotId())
	value_district := fmt.Sprintf("%d",plot.DistrictId())

	M.redis.Do(*M.ctx, radix.Cmd(nil,"MSET",
	key_x, value_x,
	key_z, value_z,
	key_coord,value_coord))

	M.redis.Do(*M.ctx, radix.Cmd(nil,"ZADD","districtZplot",value_district,value_coord))
}


func (M *MemoryCache) CacheDistrict(district *types.District){
	key_one := fmt.Sprintf("district:%d:address",district.DistrictId())
	key_two := fmt.Sprintf("district:%d:name",district.DistrictId())
	key_three := fmt.Sprintf("name:%s:district",district.StringName())
	M.redis.Do(*M.ctx, radix.FlatCmd(nil,"MSET",
	key_one,district.OwnerAddress(),
	key_two, district.StringName(),
	key_three, district.DistrictId(),
))
}

func (M *MemoryCache) CacheBlockNumber(blockNumber uint64) (error) {
	return M.redis.Do(*M.ctx,radix.FlatCmd(nil,"SET","reader_last_block",blockNumber))
}

func (M *MemoryCache) GetBlockNumber(bn *uint64) (error){
	return M.redis.Do(*M.ctx,radix.Cmd(bn,"GET","reader_last_block"))
}
