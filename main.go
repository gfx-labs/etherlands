package main

import (
	"fmt"
	"log"

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
	etherlands.save()
}

type EtherlandsContext struct {
	chain_data *DistrictConnection

	cache *MemoryCache

	plots map[uint64]*types.Plot
	best_plot uint64
}


func (E *EtherlandsContext) load() (error) {
	E.plots = make(map[uint64]*types.Plot)
	total, err := E.chain_data.GetTotalPlots()
	if(err != nil) {
		return err
	}
	var i uint64;
	for i = 1; i <= total; i++ {
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
			E.cache.CachePlot(plot)
	}
	}
	return nil
}



func (E *EtherlandsContext) save() {
	for k, v := range E.plots {
		log.Println(k,v)
		if(v != nil){
			err := v.Save();
			if err != nil {
				log.Println("failed to save plot",k)
			}
		}
	}
}
