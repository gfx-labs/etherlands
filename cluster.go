package main

import (
	"math"

	types "github.com/gfx-labs/etherlands/types"
)

type plot_id = uint64

func Cluster(W *types.World, plot_ids []uint64) [][]plot_id {
	result := make(map[plot_id]uint64)
	db := make(map[plot_id]*types.Plot)
	for _, v := range plot_ids {
		plot, err := W.GetPlot(v)
		if err == nil {
			db[v] = plot
		}
	}
	current_cluster := uint64(0)
	for k := range db {
		if _, ok := result[k]; !ok {
			find_neighbors(W, &result, db, k, current_cluster)
			current_cluster = current_cluster + 1
		}
	}
	output := make([][]plot_id, current_cluster)
	for i := range output {
		output[i] = make([]uint64, 0)
	}
	for plot, cluster := range result {
		output[cluster] = append(output[cluster], plot)
	}
	return output
}

type ClusterMetadata struct {
	OriginX int64 `json:"origin_x"`
	OriginZ int64 `json:"origin_z"`
	LengthX int64 `json:"length_x"`
	LengthZ int64 `json:"length_z"`

	Offsets [][2]int64 `json:"offsets"`
	PlotIds []uint64   `json:"plot_ids"`
}

func GenerateClusterMetadata(
	W *types.World,
	plot_ids []uint64,
) []ClusterMetadata {
	output := []ClusterMetadata{}
	clustered := Cluster(W, plot_ids)
	for _, cluster := range clustered {

		var min_x int64 = math.MaxInt64
		var max_x int64 = math.MinInt64

		var min_z int64 = math.MaxInt64
		var max_z int64 = math.MinInt64

		for _, plot_id := range cluster {
			plot, err := W.GetPlot(plot_id)
			if err == nil {
				if min_x > plot.X() {
					min_x = plot.X()
				}
				if max_x < plot.X() {
					max_x = plot.X()
				}
				if min_z > plot.Z() {
					min_z = plot.Z()
				}
				if max_z < plot.Z() {
					max_z = plot.Z()
				}
			}
		}
		offsets := [][2]int64{}
		ids := []uint64{}
		for _, plot_id := range cluster {
			plot, err := W.GetPlot(plot_id)
			if err == nil {
				offsets = append(offsets, [2]int64{plot.X() - min_x, plot.Z() - min_z})
				ids = append(ids, plot.PlotId())
			}
		}
		output = append(output, ClusterMetadata{
			OriginX: min_x,
			OriginZ: min_z,
			LengthX: 1 + max_x - min_x,
			LengthZ: 1 + max_z - min_z,
			Offsets: offsets,
			PlotIds: ids,
		})
	}
	return output
}

func find_neighbors(
	W *types.World,
	clusters *map[plot_id]uint64,
	plotdb map[plot_id]*types.Plot,
	plot uint64,
	current_cluster uint64,
) {
	var radius int64 = 1
	var jobs []uint64 = make([]uint64, 0)

	if oplot, ok := plotdb[plot]; ok {
		origin_x := oplot.X()
		origin_z := oplot.Z()
		for idx := -radius; idx <= radius; idx++ {
			for idz := -radius; idz <= radius; idz++ {
				if id, ok := W.CheckPlot(origin_x+idx, origin_z+idz); ok {
					if _, ok := (*clusters)[id]; !ok {
						if id != 0 {
							jobs = append(jobs, id)
							(*clusters)[id] = current_cluster
						}
					}
				}
			}
		}
	}
	for _, id := range jobs {
		find_neighbors(W, clusters, plotdb, id, current_cluster)
	}
	if len(jobs) == 0 {
		(*clusters)[plot] = current_cluster
	}
}
