package main

import (
	"log"
	"time"

	"github.com/gfx-labs/etherlands/types"
)

func main() {

	world := types.NewWorld()

	conn, err := NewDistrictConnection(world)
	if err != nil {
		log.Fatal("Failed to connect to District Contract:", err)
		return
	}

	total_districts, err := conn.GetTotalDistricts()
	if err != nil {
		log.Println("Could not query for total districts. Aborting")
		return
	}
	total_plots, err := conn.GetTotalPlots()
	if err != nil {
		log.Println("Could not query for total plots. Aborting")
		return
	}

	log.Println(total_districts, total_plots)

	err = world.LoadWorld(total_districts, total_plots)
	if err != nil {
		log.Println("error loading world:", err)
		return
	}

	go conn.process_events()
	go conn.start_query(5 * time.Second)

	StartWorldWeb(world)
	wz, err := StartWorldZmq(world)
	if err == nil {
		StartPrompt(world, wz)
	}
}
