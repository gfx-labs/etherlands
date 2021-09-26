package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/gfx-labs/etherlands/types"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "districts", Description: "District Summary"},
		{Text: "plots", Description: "Plot Summary"},
		{Text: "district", Description: "District Info"},
		{Text: "plot", Description: "Plot Info"},
		{Text: "stop", Description: "stops the server"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func StartPrompt(W *types.World) {
	for {
		dd := prompt.Input(">> ", completer,
			prompt.OptionTitle("Etherlands World Browser"),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionSuggestionBGColor(prompt.DarkGray),
			prompt.OptionPreviewSuggestionTextColor(prompt.Turquoise),
		)
		blocks := strings.Split(dd, " ")
		switch blocks[0] {
		case "districts":
			log.Printf("districts: %d", W.DistrictCount())
			districts := W.Districts()
			for i := 0; i < len(districts); i++ {
				fmt.Printf(
					"  %d > %s %s\n",
					districts[i].DistrictId(),
					districts[i].OwnerAddress(),
					districts[i].StringName(),
				)
			}
		case "plots":
			log.Printf("plots: %d", W.PlotCount())
		case "plot":
			if len(blocks) > 1 {
				num, err := strconv.ParseUint(blocks[1], 10, 64)
				if err == nil {
					plot, err := W.GetPlot(num)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Printf(
						"plot %d: Coordinates:(%d, %d) District: %d",
						plot.PlotId(),
						plot.X(),
						plot.Z(),
						plot.DistrictId(),
					)
				}
			}
		case "district":
			if len(blocks) > 1 {
				num, err := strconv.ParseUint(blocks[1], 10, 64)
				if err == nil {
					district, err := W.GetDistrict(num)
					if err != nil {
						log.Println(err)
						continue
					}
					log.Printf(
						"district %d: Owner: %s StringName: %s Plots:",
						district.DistrictId(),
						district.OwnerAddress(),
						district.StringName(),
					)
					plotmap := district.Plots()
					for _, v := range plotmap {
						fmt.Printf("%d:%d,%d ", v.PlotId(), v.X(), v.Z())
					}
					fmt.Print("\n")
				}
			}
		case "stop":
			return
		default:
			fmt.Println("Command not found", blocks[0])
		}
	}
}
