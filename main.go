package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gfx-labs/etherlands/types"
)



func main() {

	conn, err := NewDistrictConnection()
	if err != nil {
		log.Fatal("Failed to connect to District Contract:", err)
	}


	cache, err := NewMemoryCache(conn.ctx)
	if err != nil{
		log.Fatal("failed to connect to redis", err)
	}


	etherlands := EtherlandsContext{chain_data:conn, cache:cache}
	etherlands.load()
	go etherlands.process_events()
	go etherlands.start_events()

	fmt.Scanln();
	etherlands.save()
}

type EtherlandsContext struct {
	chain_data *DistrictConnection

	cache *MemoryCache

	plots_lock sync.RWMutex
	plots map[uint64]*types.Plot

	districts_lock sync.RWMutex
	districts map[uint64]*types.District

	best_plot uint64
	best_district uint64
}

func (E *EtherlandsContext) GetPlot(id uint64) (*types.Plot) {
	E.plots_lock.RLock()
	defer E.plots_lock.RUnlock()
	if value, ok := E.plots[id]; ok {
		return value
	}
	return nil
}
func (E *EtherlandsContext) SetPlot(plot *types.Plot){
	E.plots_lock.Lock()
	defer E.plots_lock.Unlock()
	if plot != nil {
		E.plots[plot.PlotId()] = plot
	}
}

func (E *EtherlandsContext) GetDistrict(id uint64) (*types.District) {
	E.districts_lock.RLock()
	defer E.districts_lock.RUnlock()
	if value, ok := E.districts[id]; ok {
		return value
	}
	return nil
}
func (E *EtherlandsContext) SetDistrict(district *types.District){
	E.districts_lock.Lock()
	defer E.districts_lock.Unlock()
	if district != nil {
		E.districts[district.DistrictId()] = district
	}
}









func (E *EtherlandsContext) load() (error) {
	E.plots = make(map[uint64]*types.Plot)
	E.districts = make(map[uint64]*types.District)

	var block_number uint64
	E.cache.GetBlockNumber(&block_number);
	if(E.chain_data.best_block < block_number){
		E.chain_data.best_block = block_number;
	}

	total_plots, err := E.chain_data.GetTotalPlots()
	if(err != nil) {
		return err
	}
	var i uint64;
	for i = 1; i <= total_plots; i++ {
		plot, err := types.LoadPlot(uint64(i))
		if err != nil || plot == nil{
			log.Println(fmt.Sprintf("Did not find plot %d in storage, querying chain",i))
			plot, err = E.chain_data.GetPlotInfo(i)
			if err != nil {
				log.Println("Did not find information for plot",i,"on chain")
			}else{
				log.Println("saving", plot)
				plot.Save();
			}
		}
		E.plots[i] = plot
		if(plot != nil){
			go E.cache.CachePlot(plot)
			E.best_plot = i
		}
	}

	total_districts, err := E.chain_data.GetTotalDistricts()
	if(err != nil) {
		return err
	}
	for i = 1; i <= total_districts; i++ {
		district, err := types.LoadDistrict(uint64(i))
		if err != nil || district == nil{
			log.Println(fmt.Sprintf("Did not find district %d in storage, querying chain",i))
			district, err = E.chain_data.GetDistrictInfo(i)
			if err != nil {
				log.Println("Did not find information for district",i,"on chain")
			}else{
				district.Save();
			}
		}
		E.districts[i] = district
		if(district != nil){
			go E.cache.CacheDistrict(district)
			E.best_district = i
		}
	}

	return nil
}

func (E *EtherlandsContext) save() {
	err := E.cache.CacheBlockNumber(E.chain_data.best_block)
	if(err != nil) {
		log.Println(err)
	}
	for k, v := range E.plots {
		if(v != nil){
			err := v.Save();
			if err != nil {
				log.Println("failed to save plot",k,err)
			}
		}
	}
	for k, v := range E.districts {
		if(v != nil){
			err := v.Save();
			if err != nil {
				log.Println("failed to save district",k,err)
			}
		}
	}
}

func (E *EtherlandsContext) process_events() {
	for{
		select{
		case transfer_event :=<-E.chain_data.TransferEventChannel:
			district := E.GetDistrict(transfer_event.district_id)
			if district == nil{
				district, err := E.chain_data.GetDistrictInfo(transfer_event.district_id)
				if(err == nil){
					E.SetDistrict(district)
				}
			}else{
				log.Println("updating district ", district.DistrictId())
				E.chain_data.UpdateDistrictOwner(district);
			}
		case plot_transfer_event :=<-E.chain_data.PlotTransferEventChannel:
			plot := E.GetPlot(plot_transfer_event.plot_id)
			if plot == nil {
				plot, err := E.chain_data.GetPlotInfo(plot_transfer_event.plot_id)
				if err == nil{
					E.SetPlot(plot)
				}
				}else{
				log.Println("updating plot ", plot.PlotId())
				E.chain_data.UpdatePlotDistrict(plot)
			}
		case plot_creation_event :=<-E.chain_data.PlotCreationEventChannel:
			plot, err := E.chain_data.GetPlotInfo(plot_creation_event.plot_id)
			if err != nil{
				E.SetPlot(plot)
			}
		}
	}
}

func (E* EtherlandsContext) start_events() {
	query_event_timer := start_repeating(5000)
	for{
		select {
		case _ =<-query_event_timer:
			go (func(){
				log.Println("querying block",E.chain_data.best_block)
				block, err := E.chain_data.QueryRecentEvents()
				if err != nil{
					log.Println(err)
				}else{
					E.cache.CacheBlockNumber(block)
				}
			})()
		}
	}
}
