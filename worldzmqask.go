package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	proto "github.com/gfx-labs/etherlands/proto"
	"github.com/google/uuid"
)

func (Z *WorldZmq) ask_scope(args VarArgs) {
	scope, err := args.MustGet(0)
	if Z.checkError(args, err) {
		return
	}

	switch scope {
	case "world":
		Z.ask_world_type(args)
	default:
		Z.checkError(args, errors.New("Unspecified Scope: "+scope))
	}
}

func (Z *WorldZmq) ask_world_type(args VarArgs) {
	dtype, err := args.MustGet(1)
	if Z.checkError(args, err) {
		return
	}
	switch dtype {
	case "gamer":
		Z.ask_world_gamer_field(args)
	case "plot":
		Z.ask_world_plot_field(args)
	case "town":
		Z.ask_world_town_field(args)
	case "district":
		Z.ask_world_district_field(args)
	case "links":
		addr, err := args.MustGet(2)
		if Z.checkError(args, err) {
			return
		}
		gamer_str, err := Z.W.Cache().GetLink(strings.ToLower(addr))
		if Z.checkError(args, err) {
			return
		}
		Z.sendResponse(args, gamer_str)
	case "query":
		Z.ask_world_query_field(args)
	default:
		Z.checkError(args, errors.New("Unspecified Type: "+dtype))
	}
}

func (Z *WorldZmq) ask_world_query_field(args VarArgs) {
	field, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "plot_coord":
		coord_str, err := args.MustGet(3)
		if Z.checkError(args, err) {
			return
		}
		split := strings.Split(coord_str, "_")
		if len(split) != 2 {
			Z.genericError(args, "invalid coordinate input")
			return
		}
		x, err := strconv.ParseInt(split[0], 10, 64)
		if Z.checkError(args, err) {
			return
		}
		z, err := strconv.ParseInt(split[1], 10, 64)
		if Z.checkError(args, err) {
			return
		}
		plot, err := Z.W.SearchPlot(x, z)
		if Z.checkError(args, err) {
			return
		}
		Z.sendResponse(args, strconv.FormatUint(plot.PlotId(), 10))
	case "district_by_name":
		district_name, err := args.MustGet(3)
		if Z.checkError(args, err) {
			return
		}
		district_id, err := Z.W.Cache().GetDistrictByName(district_name)
		if Z.checkError(args, err) {
			return
		}
		Z.sendResponse(args, strconv.FormatUint(district_id, 10))
	case "district_names":
		out := []string{}
		for _, v := range Z.W.Districts() {
			out = append(out, v.StringName())
		}
		Z.sendResponse(args, strings.Join(out, "_"))
	case "district_ids":
		out := []string{}
		for _, v := range Z.W.Districts() {
			out = append(out, strconv.FormatUint(v.DistrictId(), 10))
		}
		Z.sendResponse(args, strings.Join(out, "_"))
	default:
		Z.genericError(args, field)
	}
}

func (Z *WorldZmq) ask_world_town_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	town_id, err := args.MustGet(2)
	town, err := Z.W.GetTown(town_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "owner":
		Z.sendResponse(args, town.Owner().String())
	case "members":
		Z.sendResponse(args, FlattenUUIDSet(town.Members()))
	case "teams":
		string_set := make(map[string]struct{})
		string_set["member"] = struct{}{}
		string_set["outsider"] = struct{}{}
		string_set["manager"] = struct{}{}
		for _, v := range town.Teams() {
			string_set[v.Name()] = struct{}{}
		}
		Z.sendResponse(args, FlattenStringSet(string_set))
	case "team":
		Z.ask_world_town_team_field(args)

	default:
		Z.genericError(args, field)
	}
}

func (Z *WorldZmq) ask_world_town_team_field(args VarArgs) {
	field, err := args.MustGet(5)
	if Z.checkError(args, err) {
		return
	}
	town_id, err := args.MustGet(2)
	town, err := Z.W.GetTown(town_id)
	if Z.checkError(args, err) {
		return
	}
	team_id, err := args.MustGet(4)
	if Z.checkError(args, err) {
		return
	}
	team := town.Team(team_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "members":
		Z.sendResponse(args, FlattenUUIDSet(team.Members()))
	case "priority":
		Z.sendResponse(args, strconv.FormatInt(team.Priority(), 10))
	case "district":
		district_id, err := args.MustGetUint64(6)
		if Z.checkError(args, err) {
			return
		}
		flag_str, err := args.MustGet(7)
		if Z.checkError(args, err) {
			return
		}
		if flag, ok := proto.EnumValuesAccessFlag[flag_str]; !ok {
			Z.checkError(args, errors.New("Invalid flag "+flag_str))
			return
		} else {
			result_str := town.DistrictTeamPermissions().Read(district_id, team_id, flag)
			Z.sendResponse(args, strings.ToLower(result_str.String()))
		}
	default:
		Z.genericError(args, team_id)
	}

}

func (Z *WorldZmq) ask_world_plot_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	uuid_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	plot_id, err := strconv.ParseUint(uuid_str, 10, 64)
	if Z.checkError(args, err) {
		return
	}
	plot, err := Z.W.GetPlot(plot_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "x":
		Z.sendResponse(args, strconv.FormatInt(plot.X(), 10))
	case "z":
		Z.sendResponse(args, strconv.FormatInt(plot.Z(), 10))
	case "district":
		Z.sendResponse(args, strconv.FormatUint(plot.DistrictId(), 10))
	default:
		Z.genericError(args, field)
	}
}
func (Z *WorldZmq) ask_world_gamer_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	uuid_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	gamer_id, err := uuid.Parse(uuid_str)
	if Z.checkError(args, err) {
		return
	}
	gamer := Z.W.GetGamer(gamer_id)
	switch field {
	case "address":
		if gamer.Address() != "" {
			Z.sendResponse(args, gamer.Address())
		}
	case "pos":
		x, y, z := gamer.GetPosXYZ()
		Z.sendResponse(args, fmt.Sprintf("%d_%d_%d", x, y, z))
	case "town":
		Z.sendResponse(args, Z.W.TownOfGamer(gamer))
	default:
		Z.genericError(args, field)
	}
}
func (Z *WorldZmq) ask_world_district_field(args VarArgs) {
	field, err := args.MustGet(3)
	if Z.checkError(args, err) {
		return
	}
	district_str, err := args.MustGet(2)
	if Z.checkError(args, err) {
		return
	}
	district_id, err := strconv.ParseUint(district_str, 10, 64)
	if Z.checkError(args, err) {
		return
	}
	district, err := Z.W.GetDistrict(district_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	case "name":
		Z.sendResponse(args, district.StringName())
	case "plots":
		plots := Z.W.PlotsOfDistrict(district.DistrictId())
		temp := make([]string, len(plots))
		for i := 0; i < len(plots); i++ {
			temp[i] = strconv.FormatUint(plots[i], 10)
		}
		Z.sendResponse(args, strings.Join(temp, "_"))
	case "clusters":
		clusters := Z.W.Cache().GetClusters(district.DistrictId())
		value := ""
		for _, cluster := range clusters {
			value = value + fmt.Sprintf(
				"%d:%d:%d",
				cluster.OriginX,
				cluster.OriginZ,
				len(cluster.Offsets),
			)
			value = value + "@"
		}
		if len(value) > 1 {
			value = value[:len(value)-1]
		}
		Z.sendResponse(args, value)
	case "owner_addr":
		Z.sendResponse(args, district.OwnerAddress())
	default:
		Z.genericError(args, field)
	}
}
