package types

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gfx-labs/etherlands/zset"
	"github.com/google/uuid"
)

type World struct {
	plots         map[FamilyKey]*Plot
	plot_district *zset.ZSet
	plot_location map[[2]int64]uint64
	plots_lock    sync.RWMutex

	districts      map[FamilyKey]*District
	district_owner *zset.ZSetStr
	districts_lock sync.RWMutex

	gamers      map[FamilyKey]*Gamer
	gamers_lock sync.RWMutex

	towns map[FamilyKey]*Town

	DistrictRequests chan uint64
	PlotRequests     chan uint64

	DistrictIn chan DistrictChainInfo
	PlotIn     chan PlotChainInfo

	linkermap *LinkerMap

	cache *MemoryCache
}

type DistrictChainInfo struct {
	DistrictId uint64
	Nickname   [24]byte
	Owner      string
}

type PlotChainInfo struct {
	PlotId     uint64
	DistrictId uint64
	X          int64
	Z          int64
}

func NewWorld() *World {
	output := &World{
		plots:          make(map[FamilyKey]*Plot),
		districts:      make(map[FamilyKey]*District),
		gamers:         make(map[FamilyKey]*Gamer),
		towns:          make(map[FamilyKey]*Town),
		plot_location:  make(map[[2]int64]uint64),
		plot_district:  zset.CreateZSet(),
		district_owner: zset.CreateZSetStr(),

		DistrictRequests: make(chan uint64, 100),
		PlotRequests:     make(chan uint64, 100),
		DistrictIn:       make(chan DistrictChainInfo, 100),
		PlotIn:           make(chan PlotChainInfo, 100),

		linkermap: NewLinkerMap(time.Minute * 15),
	}
	memcache, err := NewMemoryCache()
	if err == nil {
		output.cache = memcache
	} else {
		log.Println("no redis - running in dumb mode")
	}
	return output
}

func (W *World) CreateLinkRequest(message string) {
	W.linkermap.Add(message)
}

func (W *World) HonorLinkRequest(gamer_id uuid.UUID, address string, message string) bool {
	if W.linkermap.Check(strings.ToLower(message)) {
		gamer := W.GetGamer(gamer_id)
		gamer.SetAddress(address)
		W.UpdateGamer(gamer)
		return true
	}
	return false
}

func (W *World) UpdateGamer(gamer *Gamer) {
	W.gamers_lock.Lock()
	defer W.gamers_lock.Unlock()
	if _, ok := W.gamers[gamer.GetKey()]; !ok {
		W.gamers[gamer.GetKey()] = gamer
	}
	W.cache.CacheGamer(W.gamers[gamer.GetKey()])
	go W.gamers[gamer.GetKey()].Save()
}

func (W *World) UpdatePlot(plot *Plot) {
	W.plots_lock.Lock()
	defer W.plots_lock.Unlock()
	W.plot_district.AddOrUpdate(plot.PlotId(), plot.DistrictId(), plot)
	W.plot_location[[2]int64{plot.X(), plot.Z()}] = plot.PlotId()
	if _, ok := W.plots[plot.GetKey()]; !ok {
		W.plots[plot.GetKey()] = plot
	}
	W.cache.CachePlot(W.plots[plot.GetKey()])
}

func (W *World) UpdateDistrict(district *District) {
	if district == nil {
		return
	}
	W.districts_lock.Lock()
	defer W.districts_lock.Unlock()
	W.district_owner.AddOrUpdate(district.DistrictId(), district.OwnerAddress(), district)
	if _, ok := W.districts[district.GetKey()]; !ok {
		W.districts[district.GetKey()] = district
	}
	W.cache.CacheDistrict(W.districts[district.GetKey()])
	go W.districts[district.GetKey()].Save()
}

// every plot is always loaded into memory
// plot location is immutable!!!
func (W *World) LoadWorld(district_count uint64, plot_count uint64) error {
	for i := uint64(1); i < plot_count; i++ {
		plot, err := W.LoadPlot(i)
		if err != nil {
			log.Println("failed to read plot", err)
			go func() {
				W.PlotRequests <- i
			}()
		} else {
			W.UpdatePlot(plot)
		}
	}
	for i := uint64(1); i < district_count; i++ {
		district, err := W.LoadDistrict(i)
		if err != nil {
			log.Println("failed to read district", err)
			go func() {
				W.DistrictRequests <- i
			}()
		} else {
			W.UpdateDistrict(district)
		}
	}
	files, err := ListStruct("gamers")
	if err == nil {
		for i := 0; i < len(files); i++ {
			gamer_id, err := uuid.Parse(files[i].Name())
			if err == nil {
				gamer, err := W.LoadGamer(gamer_id)
				if err != nil {
					log.Println("failed to read gamer", err)
				} else {
					W.UpdateGamer(gamer)
				}
			}
		}
	}
	// don't start listening for requests until after we load in from memory
	go func() {
		for {
			district_info := <-W.DistrictIn
			log.Println("district info:", district_info)
			district, _ := W.GetDistrict(district_info.DistrictId)
			if district == nil {
				W.NewDistrict(
					district_info.DistrictId,
					district_info.Owner,
					district_info.Nickname,
				)
			} else {
				district.SetNickname(district_info.Nickname)
				district.SetOwnerAddress(district_info.Owner)
			}
			W.UpdateDistrict(district)
		}
	}()

	go func() {
		for {
			plot_info := <-W.PlotIn
			log.Println("plot info:", plot_info)
			plot, err := W.GetPlot(plot_info.PlotId)
			if err == nil {
				plot = W.newPlot(
					plot_info.X,
					plot_info.Z,
					plot_info.PlotId,
					plot_info.DistrictId,
				)
			} else {
				plot.SetDistrictId(plot_info.DistrictId)
			}
			W.UpdatePlot(plot)
		}
	}()

	return nil
}

func (W *World) SaveGamer(gamer *Gamer) {
	// save gamer
	go gamer.Save()
}
