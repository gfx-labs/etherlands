package main

import (
	types "github.com/gfx-labs/etherlands/types"
)

type plot_id = uint64;

func (E *EtherlandsContext) Cluster(plot_ids []uint64) [][]plot_id{
	result := make(map[plot_id]uint64);
	db := make(map[plot_id]*types.Plot);
	for _, v := range plot_ids {
		db[v] = E.GetPlot(v)
	}
	current_cluster := uint64(0)
	for k := range db {
		if _, ok := result[k]; !ok{
			find_neighbors(&result,db,E.plot_location,k,current_cluster);
			current_cluster = current_cluster + 1;
		}
	}
	output := make([][]plot_id,current_cluster)
	for i := range output{
		output[i] = make([]uint64,0)
	}
	for plot, cluster := range result {
		output[cluster] = append(output[cluster], plot)
	}
	return output
}

type ClusterMetadata struct {
	OriginX int64 `json:"origin_x"`
	OriginZ int64 `json:"origin_z"`
}

func find_neighbors(
	clusters *map[plot_id]uint64,
	plotdb map[plot_id]*types.Plot,
	locationdb map[[2]int64]uint64,
	plot uint64,
	current_cluster uint64,
){
	var radius int64 = 1;
	var jobs []uint64 = make([]uint64,0);

	if oplot, ok := plotdb[plot]; ok {
		origin_x := oplot.X()
		origin_z := oplot.Z()
		for idx := -radius; idx <= radius; idx++ {
			for idz := -radius; idz <= radius; idz++ {
				loc := [2]int64{
					origin_x + idx,
					origin_z + idz,
				}
				if id, ok := locationdb[loc]; ok{
					if _, ok := plotdb[id]; ok{
						if _, ok := (*clusters)[id]; !ok {
							if(id != 0){
								jobs = append(jobs, id);
								(*clusters)[id] = current_cluster;
							}
						}
					}
				}
			}
		}
	}
	for _, id := range jobs {
		find_neighbors(clusters,plotdb,locationdb,id,current_cluster)
	}
	if len(jobs) == 0 {
		(*clusters)[plot] = current_cluster;
	}
}
