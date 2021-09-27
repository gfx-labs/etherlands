package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	types "github.com/gfx-labs/etherlands/types"
	"github.com/google/uuid"
	"github.com/zeromq/goczmq"
)

type WorldZmq struct {
	W *types.World

	publisher  *goczmq.Channeler
	subscriber *goczmq.Channeler
}

type VarArgs []string

func StartWorldZmq(world *types.World) error {
	publisher := goczmq.NewPubChanneler("tcp://*:10105")
	subscriber := goczmq.NewSubChanneler("tcp://127.0.0.1:10106", "GET")
	zmq := &WorldZmq{W: world, publisher: publisher, subscriber: subscriber}
	go zmq.StartPublishing()
	go zmq.StartListening()
	return nil
}

func (Z *WorldZmq) StartPublishing() {
}

func (Z *WorldZmq) StartListening() {
	log.Println("now listening")
	for {
		var args VarArgs
		message := <-Z.subscriber.RecvChan
		verb := string(message[0])
		args = strings.Split(string(message[1]), ":")
		log.Printf("[%s] %s\n", verb, args.Command())
		switch verb {
		case "GET":
			Z.get_scope(args)
		default:
			log.Println("Unrecognized Verb:", verb)
		}
	}
}

func (Z *WorldZmq) get_scope(args VarArgs) {
	scope, err := args.MustGet(0)
	if Z.checkError(args, err) {
		return
	}

	switch scope {
	case "world":
		Z.get_world_type(args)
	default:
		Z.checkError(args, errors.New("Unspecified Scope: "+scope))
	}
}

func (Z *WorldZmq) get_world_type(args VarArgs) {
	dtype, err := args.MustGet(1)
	if Z.checkError(args, err) {
		return
	}
	switch dtype {
	case "gamer":
		Z.get_world_gamer_field(args)
	case "plot":
		Z.get_world_plot_field(args)
	case "district":
		Z.get_world_plot_field(args)
	case "links":
		addr, err := args.MustGet(2)
		if Z.checkError(args, err) {
			return
		}
		gamer_str, err := Z.W.Cache().GetLink(addr)
		if Z.checkError(args, err) {
			return
		}
		Z.sendResponse(args, gamer_str)
	case "query":
		Z.get_world_query_field(args)
	default:
		Z.checkError(args, errors.New("Unspecified Type: "+dtype))
	}
}

func (Z *WorldZmq) get_world_query_field(args VarArgs) {
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

func (Z *WorldZmq) get_world_plot_field(args VarArgs) {
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
		Z.checkError(args, errors.New("Unspecified Field: "+field))
	}
}
func (Z *WorldZmq) get_world_gamer_field(args VarArgs) {
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
	default:
		Z.checkError(args, errors.New("Unspecified Field: "+field))
	}
}
func (Z *WorldZmq) get_world_district_field(args VarArgs) {
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
	case "owner_addr":
		Z.sendResponse(args, district.OwnerAddress())
	default:
		Z.checkError(args, errors.New("Unspecified Field: "+field))
	}
}

func (Z *WorldZmq) sendResponse(args VarArgs, content string) {
	Z.publisher.SendChan <- [][]byte{
		[]byte(args.Command()),
		[]byte(content),
	}
}

func (Z *WorldZmq) checkError(args VarArgs, err error) bool {
	if err != nil {
		Z.publisher.SendChan <- [][]byte{
			[]byte(args.Command()),
			[]byte("error:" + err.Error()),
		}
		return true
	}
	return false
}

func (Z *WorldZmq) genericError(args VarArgs, offender string) bool {
	return Z.checkError(args, errors.New(args.Command()+": "+offender))
}

func (Args *VarArgs) Command() string {
	return strings.Join(*Args, ":")
}

func (Args *VarArgs) MustGet(idx int) (string, error) {
	if len(*Args) > idx {
		return (*Args)[idx], nil
	}
	return "", errors.New("Variable out of bounds")
}

func (Args *VarArgs) MightGet(idx int) (string, bool) {
	if len(*Args) > idx {
		return (*Args)[idx], true
	}
	return "", false
}
