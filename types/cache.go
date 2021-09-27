package types

import (
	"errors"
	"sync"
	"time"

	utils "github.com/gfx-labs/etherlands/utils"
)

type WorldCache struct {
	links      map[string]string
	links_lock sync.RWMutex

	name_district      map[string]uint64
	name_district_lock sync.RWMutex

	clusters      map[uint64][]ClusterMetadata
	cluster_lock  sync.RWMutex
	cluster_limit *utils.RateLimit

	W *World
}

func (W *World) NewWorldCache() (*WorldCache, error) {
	return &WorldCache{
		W:             W,
		links:         make(map[string]string),
		name_district: make(map[string]uint64),
		clusters:      make(map[uint64][]ClusterMetadata),
		cluster_limit: utils.NewRateLimiter(1 * time.Minute),
	}, nil
}

func (M *WorldCache) CachePlot(plot *Plot) {
}

func (M *WorldCache) CacheTown(town *Town) {
}

func (M *WorldCache) CacheDistrict(district *District) {
	M.name_district_lock.Lock()
	M.name_district[district.StringName()] = district.DistrictId()
	M.name_district_lock.Unlock()

}

func (M *WorldCache) CacheGamer(gamer *Gamer) {
	M.links_lock.Lock()
	M.links[gamer.MinecraftId().String()] = gamer.Address()
	M.links[gamer.Address()] = gamer.MinecraftId().String()
	M.links_lock.Unlock()
}

func (M *WorldCache) GetDistrictByName(input string) (uint64, error) {
	M.name_district_lock.RLock()
	defer M.name_district_lock.RUnlock()
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
