package types

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/mediocregopher/radix/v4"
)

type MemoryCache struct {
	redis radix.Client
	ctx   *context.Context

	links      map[string]string
	links_lock sync.RWMutex

	name_district      map[string]uint64
	name_district_lock sync.RWMutex
}

func NewMemoryCache() (*MemoryCache, error) {
	ctx := context.Background()
	redis, err := (radix.PoolConfig{}).New(ctx, "tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, err
	}
	return &MemoryCache{redis: redis, ctx: &ctx,
		links:         make(map[string]string),
		name_district: make(map[string]uint64),
	}, nil
}

func (M *MemoryCache) CachePlot(plot *Plot) {
	key_x := fmt.Sprintf("plot:%d:x", plot.PlotId())
	value_x := fmt.Sprintf("%d", plot.X())
	key_z := fmt.Sprintf("plot:%d:z", plot.PlotId())
	value_z := fmt.Sprintf("%d", plot.Z())
	key_coord := fmt.Sprintf("plot_coord:%d_%d", plot.X(), plot.Z())
	value_coord := fmt.Sprintf("%d", plot.PlotId())
	//key_district := fmt.Sprintf("plot:%d:district",plot.PlotId())
	value_district := fmt.Sprintf("%d", plot.DistrictId())

	M.redis.Do(*M.ctx, radix.Cmd(nil, "MSET",
		key_x, value_x,
		key_z, value_z,
		key_coord, value_coord))

	M.redis.Do(*M.ctx, radix.Cmd(nil, "ZADD", "districtZplot", value_district, value_coord))
}

func (M *MemoryCache) CacheDistrict(district *District) {
	M.name_district_lock.Lock()
	M.name_district[district.StringName()] = district.DistrictId()
	M.name_district_lock.Unlock()

	key_one := fmt.Sprintf("district:%d:address", district.DistrictId())
	M.redis.Do(*M.ctx, radix.FlatCmd(nil, "MSET",
		key_one, district.OwnerAddress(),
	))
	M.redis.Do(*M.ctx, radix.FlatCmd(nil, "HSET",
		"name_district", district.StringName(), district.DistrictId(),
	))
	M.redis.Do(*M.ctx, radix.FlatCmd(nil, "HSET",
		"district_name", district.DistrictId(), district.StringName(),
	))
}

func (M *MemoryCache) GetDistrictByName(input string) (uint64, error) {
	M.name_district_lock.RLock()
	defer M.name_district_lock.RUnlock()
	if v, ok := M.name_district[input]; ok {
		return v, nil
	}
	return 0, errors.New("no district found")
}
func (M *MemoryCache) GetLink(input string) (string, error) {
	M.links_lock.RLock()
	defer M.links_lock.RUnlock()
	if v, ok := M.links[input]; ok {
		return v, nil
	}
	return "", errors.New("no link found")
}

func (M *MemoryCache) CacheGamer(gamer *Gamer) {
	M.links_lock.Lock()
	M.links[gamer.MinecraftId().String()] = gamer.Address()
	M.links[gamer.Address()] = gamer.MinecraftId().String()
	M.links_lock.Unlock()
	M.redis.Do(*M.ctx, radix.FlatCmd(nil, "HMSET", "gamer_links",
		gamer.MinecraftId().String(), gamer.Address(),
		gamer.Address(), gamer.MinecraftId().String(),
	))

}
