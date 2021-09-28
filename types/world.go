package types

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type World struct {
	plots      map[FamilyKey]*Plot
	plots_lock sync.RWMutex

	districts      map[FamilyKey]*District
	districts_lock sync.RWMutex

	gamers      map[FamilyKey]*Gamer
	gamers_lock sync.RWMutex

	towns      map[FamilyKey]*Town
	towns_lock sync.RWMutex

	DistrictRequests chan uint64
	PlotRequests     chan uint64

	DistrictIn chan DistrictChainInfo
	PlotIn     chan PlotChainInfo

	sendChan chan [2]string

	linkermap *LinkerMap

	cache *WorldCache
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
		plots:     make(map[FamilyKey]*Plot),
		districts: make(map[FamilyKey]*District),
		gamers:    make(map[FamilyKey]*Gamer),
		towns:     make(map[FamilyKey]*Town),

		DistrictRequests: make(chan uint64, 100),
		PlotRequests:     make(chan uint64, 100),
		DistrictIn:       make(chan DistrictChainInfo, 100),
		PlotIn:           make(chan PlotChainInfo, 100),

		linkermap: NewLinkerMap(time.Minute * 15),
	}
	memcache, err := output.NewWorldCache()
	if err == nil {
		output.cache = memcache
	} else {
		log.Println("failed to init cache")
	}
	return output
}

func (W *World) Cache() *WorldCache {
	return W.cache
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
	gamer.Update()
}
func (G *Gamer) Update() {
	if G.GetKey().datatype == GAMER_FAMILY {
		G.W.gamers_lock.Lock()
		defer G.W.gamers_lock.Unlock()
		if _, ok := G.W.gamers[G.GetKey()]; !ok {
			G.W.gamers[G.GetKey()] = G
		}
		G.W.cache.CacheGamer(G.W.gamers[G.GetKey()])
		go G.W.gamers[G.GetKey()].Save()
	}
}

func (W *World) UpdatePlot(plot *Plot) {
	plot.Update()
}
func (plot *Plot) Update() {
	plot.W.plots_lock.Lock()
	defer plot.W.plots_lock.Unlock()
	if _, ok := plot.W.plots[plot.GetKey()]; !ok {
		plot.W.plots[plot.GetKey()] = plot
	}
	plot.W.cache.CachePlot(plot.W.plots[plot.GetKey()])
}

func (W *World) UpdateTown(town *Town) {
	W.towns_lock.Lock()
	defer W.towns_lock.Unlock()
	if _, ok := W.towns[town.GetKey()]; !ok {
		W.towns[town.GetKey()] = town
	}
	W.cache.CacheTown(W.towns[town.GetKey()])
	go W.towns[town.GetKey()].Save()
}

func (W *World) DeleteTown(town *Town) {
	W.towns_lock.Lock()
	defer W.towns_lock.Unlock()
	if _, ok := W.towns[town.GetKey()]; !ok {
		return
	}
	DeleteStruct("towns", town.Name())
	W.cache.DeleteTown(W.towns[town.GetKey()])
	delete(W.towns, town.GetKey())
}

func (W *World) UpdateDistrict(district *District) {
	if district == nil {
		return
	}
	W.districts_lock.Lock()
	defer W.districts_lock.Unlock()
	if _, ok := W.districts[district.GetKey()]; !ok {
		W.districts[district.GetKey()] = district
	}
	go W.cache.CacheDistrict(W.districts[district.GetKey()])
	go W.districts[district.GetKey()].Save()
}

// every plot is always loaded into memory
// plot location is immutable!!!
func (W *World) LoadWorld(district_count uint64, plot_count uint64) error {
	for i := uint64(1); i <= plot_count; i++ {
		plot, err := W.LoadPlot(i)
		if err != nil {
			log.Println("failed to read plot", err)
			go func(j uint64) {
				W.PlotRequests <- j
			}(i)
		} else {
			W.UpdatePlot(plot)
		}
	}
	for i := uint64(1); i <= district_count; i++ {
		district, err := W.LoadDistrict(i)
		if err != nil {
			log.Println("failed to read district", err)
			go func(j uint64) {
				W.DistrictRequests <- j
			}(i)
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
	town_files, err := ListStruct("towns")
	for i := 0; i < len(town_files); i++ {
		town_id := town_files[i].Name()
		town, err := W.LoadTown(town_id)
		if err != nil {
			log.Println("failed to read town", err)
		} else {
			W.UpdateTown(town)
		}
	}
	// don't start listening for requests until after we load in from memory
	go func() {
		for {
			district_info := <-W.DistrictIn
			log.Println("district info:", district_info)
			district, err := W.GetDistrict(district_info.DistrictId)
			if err != nil {
				W.NewDistrict(
					district_info.DistrictId,
					district_info.Owner,
					district_info.Nickname,
				)
			} else {
				if *district.Nickname() != district_info.Nickname {
					district.SetNickname(district_info.Nickname)
				}
				if district.OwnerAddress() != district_info.Owner {
					district.SetOwnerAddress(district_info.Owner)
				}
			}
			W.UpdateDistrict(district)
		}
	}()

	go func() {
		for {
			plot_info := <-W.PlotIn
			log.Println("plot info:", plot_info)
			plot, err := W.GetPlot(plot_info.PlotId)
			if err != nil {
				plot = W.newPlot(
					plot_info.X,
					plot_info.Z,
					plot_info.PlotId,
					plot_info.DistrictId,
				)
			} else {
				if plot.DistrictId() != plot_info.DistrictId {
					plot.SetDistrictId(plot_info.DistrictId)
				}
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
