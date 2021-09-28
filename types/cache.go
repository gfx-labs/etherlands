package types

import (
	"errors"
	"sync"
	"time"

	utils "github.com/gfx-labs/etherlands/utils"
	"github.com/gfx-labs/etherlands/zset"
)

type WorldCache struct {
	links      map[string]string
	links_lock sync.RWMutex

	plot_district *zset.ZSet
	plot_location map[[2]int64]uint64
	plot_lock     sync.RWMutex

	name_district  map[string]uint64
	district_owner *zset.ZSetStr
	district_town  *zset.ZSetStr
	district_lock  sync.RWMutex

	clusters      map[uint64][]ClusterMetadata
	cluster_lock  sync.RWMutex
	cluster_limit *utils.RateLimit

	uuid_town      *zset.ZSetUUIDStr
	uuid_town_lock sync.RWMutex

	W *World
}

func (W *World) NewWorldCache() (*WorldCache, error) {
	return &WorldCache{
		W:              W,
		links:          make(map[string]string),
		plot_location:  make(map[[2]int64]uint64),
		plot_district:  zset.CreateZSet(),
		name_district:  make(map[string]uint64),
		district_owner: zset.CreateZSetStr(),
		district_town:  zset.CreateZSetStr(),
		clusters:       make(map[uint64][]ClusterMetadata),
		cluster_limit:  utils.NewRateLimiter(1 * time.Minute),
		uuid_town:      zset.CreateZSetUUIDStr(),
	}, nil
}

func (M *WorldCache) CachePlot(plot *Plot) {
	M.plot_lock.Lock()
	defer M.plot_lock.Unlock()
	M.plot_district.AddOrUpdate(plot.PlotId(), plot.DistrictId(), struct{}{})
	M.plot_location[[2]int64{plot.X(), plot.Z()}] = plot.PlotId()
}

func (M *WorldCache) CacheTown(town *Town) {
	M.uuid_town_lock.Lock()
	defer M.uuid_town_lock.Unlock()
	M.uuid_town.AddOrUpdate(town.Owner(), town.Name(), town)
}

func (M *WorldCache) DeleteTown(town *Town) {
	M.uuid_town_lock.Lock()
	defer M.uuid_town_lock.Unlock()
	M.uuid_town.Remove(town.Owner())
}

func (M *WorldCache) CacheDistrict(district *District) {
	M.district_lock.Lock()
	defer M.district_lock.Unlock()
	M.name_district[district.StringName()] = district.DistrictId()
	M.district_owner.AddOrUpdate(district.DistrictId(), district.OwnerAddress(), district)
	M.district_town.AddOrUpdate(district.DistrictId(), district.Town(), district)
}

func (M *WorldCache) CacheGamer(gamer *Gamer) {
	M.links_lock.Lock()
	defer M.links_lock.Unlock()
	M.links[gamer.MinecraftId().String()] = gamer.Address()
	M.links[gamer.Address()] = gamer.MinecraftId().String()
	M.uuid_town_lock.Lock()
	M.uuid_town.AddOrUpdate(gamer.MinecraftId(), gamer.Town(), struct{}{})
	M.uuid_town_lock.Unlock()
}

func (M *WorldCache) GetDistrictByName(input string) (uint64, error) {
	M.district_lock.RLock()
	defer M.district_lock.RUnlock()
	if v, ok := M.name_district[input]; ok {
		return v, nil
	}
	return 0, errors.New("no district found")
}
func (M *WorldCache) GetLink(input string) (string, error) {
	M.links_lock.RLock()
	defer M.links_lock.RUnlock()
	if v, ok := M.links[input]; ok {
		return v, nil
	}
	return "", errors.New("no link found")
}

func (M *WorldCache) CheckPlot(x, z int64) (uint64, bool) {
	M.plot_lock.RLock()
	defer M.plot_lock.RUnlock()
	if v, ok := M.plot_location[[2]int64{x, z}]; ok {
		return v, true
	}
	return 0, false
}

func (M *WorldCache) CacheClusters(input uint64, clusters []ClusterMetadata) {

	M.cluster_lock.Lock()
	M.clusters[input] = clusters
	M.cluster_lock.Unlock()
}

func (M *WorldCache) GetClusters(input uint64) []ClusterMetadata {
	if !M.cluster_limit.Check(input) {
		M.cluster_lock.RLock()
		if v, ok := M.clusters[input]; ok {
			M.cluster_lock.RUnlock()
			return v
		}
		M.cluster_lock.RUnlock()
	}
	count := M.W.PlotsOfDistrict(input)
	clustered := M.W.GenerateClusterMetadata(count)
	M.CacheClusters(input, clustered)
	return clustered
}
