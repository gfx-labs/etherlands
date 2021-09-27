package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func (Z *WorldZmq) hit_scope(args VarArgs) {
	scope, err := args.MustGet(0)
	if Z.checkError(args, err) {
		return
	}

	switch scope {
	case "world":
		Z.hit_world_type(args)
	default:
		Z.checkError(args, errors.New("Unspecified Scope: "+scope))
	}
}

func (Z *WorldZmq) smp_link_request(message string) error {
	args := strings.Split(message, ":")
	if len(args) == 4 {
		_, err := uuid.Parse(args[0])
		if err != nil {
			return errors.New(fmt.Sprintf("malformed uuid %s", args[0]))
		}
		if args[1] != "" && args[2] != "" && args[3] != "" {
			Z.W.CreateLinkRequest(strings.ToLower(message))
			return nil
		}
	}
	return errors.New(fmt.Sprintf("invalid input %s %v", message, args))
}

func (Z *WorldZmq) hit_world_type(args VarArgs) {
	dtype, err := args.MustGet(1)
	if Z.checkError(args, err) {
		return
	}
	switch dtype {
	case "gamer":
		Z.hit_world_gamer_field(args)
	case "plot":
		Z.hit_world_plot_field(args)
	case "district":
		Z.hit_world_district_field(args)
	case "link_request":
		link_info, err := args.MustGet(2)
		if Z.checkError(args, err) {
			return
		}
		Z.smp_link_request(link_info)
		if Z.checkError(args, err) {
			return
		}
	default:
		Z.checkError(args, errors.New("Unspecified Type: "+dtype))
	}
}

func (Z *WorldZmq) hit_world_plot_field(args VarArgs) {
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
	_, err = Z.W.GetPlot(plot_id)
	if Z.checkError(args, err) {
		return
	}
	switch field {
	default:
		Z.genericError(args, field)
	}
}
func (Z *WorldZmq) hit_world_gamer_field(args VarArgs) {
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
	case "pos":
		x, err := args.MustGetInt64(4)
		if Z.checkError(args, err) {
			return
		}
		y, err := args.MustGetInt64(5)
		if Z.checkError(args, err) {
			return
		}
		z, err := args.MustGetInt64(6)
		if Z.checkError(args, err) {
			return
		}
		gamer.SetPosXYZ(x, y, z)
	default:
		Z.genericError(args, field)
	}
}

func (Z *WorldZmq) hit_world_district_field(args VarArgs) {
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
