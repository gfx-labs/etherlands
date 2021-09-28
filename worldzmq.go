package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	logger "github.com/gfx-labs/etherlands/logger"
	"github.com/google/uuid"

	types "github.com/gfx-labs/etherlands/types"
	zmq "github.com/pebbe/zmq4"
)

type WorldZmq struct {
	W *types.World

	publisher  *zmq.Socket
	subscriber *zmq.Socket

	recvChan chan [2]string
	sendChan chan [2]string

	mutexes sync.Map
}

func (Z *WorldZmq) lock(name string) func() {
	value, _ := Z.mutexes.LoadOrStore(name, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

type VarArgs []string

func StartWorldZmq(world *types.World) (*WorldZmq, error) {
	publisher, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		return nil, err
	}
	publisher.Bind("tcp://*:10105")
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}
	subscriber.Bind("tcp://127.0.0.1:10106")
	subscriber.SetSubscribe("")
	zmq := &WorldZmq{W: world, publisher: publisher, subscriber: subscriber,
		recvChan: make(chan [2]string, 100),
		sendChan: make(chan [2]string, 100),
	}
	go zmq.StartPublishing()
	go zmq.StartSubscribing()
	go zmq.StartListening()
	return zmq, nil
}

func (Z *WorldZmq) StartPublishing() {
	for {
		pair := <-Z.sendChan
		Z.publisher.Send(pair[0], zmq.SNDMORE)
		Z.publisher.Send(pair[1], 0)
	}
}

func (Z *WorldZmq) StartSubscribing() {
	for {
		verb, err1 := Z.subscriber.Recv(0)
		command, err2 := Z.subscriber.Recv(0)
		if err1 == nil && err2 == nil {
			Z.recvChan <- [2]string{verb, command}
		}
	}
}

func (Z *WorldZmq) StartListening() {
	for {
		var args VarArgs
		message := <-Z.recvChan
		verb := string(message[0])
		args = strings.Split(string(message[1]), ":")
		logger.Log.Printf("[%s] %s\n", verb, args.Command())
		switch verb {
		case "ASK":
			Z.ask_scope(args)
		case "HIT":
			Z.hit_scope(args)
		default:
			logger.Log.Println("Unrecognized Verb:", verb)
		}
	}
}

func (Z *WorldZmq) sendResponse(args VarArgs, content string) {
	Z.sendChan <- [2]string{
		args.Command(),
		string(content),
	}
	logger.Log.Printf("[OUT] %s %s\n", args.Command(), content)
}

func (Z *WorldZmq) checkGamerError(gamer *types.Gamer, err error) bool {
	return Z.checkUUIDError(gamer.MinecraftId(), err)
}

func (Z *WorldZmq) checkUUIDError(target uuid.UUID, err error) bool {
	if err != nil {
		Z.sendChan <- [2]string{
			"CHAT",
			fmt.Sprintf("gamer:%s:[Error] %s", target.String(), err.Error()),
		}
		logger.Log.Printf("[CHAT] [%s] %s\n", target.String(), err.Error())
		return true
	}
	return false
}

func (Z *WorldZmq) sendGamerResult(gamer *types.Gamer, result string) {
	Z.sendUUIDResult(gamer.MinecraftId(), result)
}
func (Z *WorldZmq) sendUUIDResult(target uuid.UUID, result string) {
	Z.sendChan <- [2]string{
		"CHAT",
		fmt.Sprintf("gamer:%s:%s", target.String(), result),
	}
	logger.Log.Printf("[CHAT] [%s] %s\n", target.String(), result)
}
func (Z *WorldZmq) sendTownResult(target string, result string) {
	Z.sendChan <- [2]string{
		"CHAT",
		fmt.Sprintf("team:%s:%s", target, result),
	}
	logger.Log.Printf("[CHAT] [%s] %s\n", target, result)
}

func (Z *WorldZmq) checkError(args VarArgs, err error) bool {
	if err != nil {
		Z.sendChan <- [2]string{
			args.Command(),
			"error:" + err.Error(),
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
	return "", errors.New("Missing Argument" + strconv.FormatInt(int64(idx), 10))
}

func (Args *VarArgs) MustGetGamer(W *types.World, idx int) (*types.Gamer, error) {
	uuid_str, err := Args.MustGet(idx)
	if err != nil {
		return nil, err
	}
	gamer_id, err := uuid.Parse(uuid_str)
	if err != nil {
		return nil, err
	}
	gamer := W.GetGamer(gamer_id)
	if err != nil {
		return nil, err
	}
	return gamer, nil
}

func (Args *VarArgs) MustGetUint64(idx int) (uint64, error) {
	if len(*Args) > idx {
		return strconv.ParseUint((*Args)[idx], 10, 64)
	}
	return 0, errors.New("Missing Argument" + strconv.FormatInt(int64(idx), 10))
}

func (Args *VarArgs) MustGetInt64(idx int) (int64, error) {
	if len(*Args) > idx {
		return strconv.ParseInt((*Args)[idx], 10, 64)
	}
	return 0, errors.New("Variable out of bounds")
}

func (Args *VarArgs) MightGet(idx int) (string, bool) {
	if len(*Args) > idx {
		return (*Args)[idx], true
	}
	return "", false
}
func FlattenStringSet(set map[string]struct{}) string {
	if len(set) == 0 {
		return ""
	}
	out := ""
	first := true
	for k := range set {
		if first {
			out = out + k
			first = false
		} else {
			out = out + ";" + k
		}
	}
	return out
}
func FlattenStringAny(set map[string]interface{}) string {
	if len(set) == 0 {
		return ""
	}
	out := ""
	first := true
	for k := range set {
		if first {
			out = out + k
			first = false
		} else {
			out = out + ";" + k
		}
	}
	return out
}

func FlattenUintSet(set map[uint64]struct{}) string {
	if len(set) == 0 {
		return ""
	}
	out := ""
	first := true
	for k := range set {
		if first {
			out = out + strconv.FormatUint(k, 10)
			first = false
		} else {
			out = out + ";" + strconv.FormatUint(k, 10)
		}
	}
	return out
}
func FlattenUintSlice(slice []uint64) string {
	if len(slice) == 0 {
		return ""
	}
	out := ""
	first := true
	for _, k := range slice {
		if first {
			out = out + strconv.FormatUint(k, 10)
			first = false
		} else {
			out = out + ";" + strconv.FormatUint(k, 10)
		}
	}
	return out
}

func FlattenUUIDSet(set map[uuid.UUID]struct{}) string {
	if len(set) == 0 {
		return ""
	}
	out := ""
	first := true
	for k := range set {
		if first {
			out = out + k.String()
			first = false
		} else {
			out = out + ";" + k.String()
		}
	}
	return out
}
