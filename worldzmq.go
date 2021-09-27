package main

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	logger "github.com/gfx-labs/etherlands/logger"

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
	return "", errors.New("Variable out of bounds")
}

func (Args *VarArgs) MustGetUint64(idx int) (uint64, error) {
	if len(*Args) > idx {
		return strconv.ParseUint((*Args)[idx], 10, 64)
	}
	return 0, errors.New("Variable out of bounds")
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
