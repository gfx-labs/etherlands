package types

import (
	"sync"
	"time"
)

type LinkerMap struct {
	internal map[string]struct{}
	timeout  time.Duration
	sync.Mutex
}

func NewLinkerMap(duration time.Duration) *LinkerMap {
	return &LinkerMap{
		internal: make(map[string]struct{}),
		timeout:  duration,
	}
}

func (L *LinkerMap) Check(keyword string) bool {
	L.Lock()
	defer L.Unlock()
	if _, ok := L.internal[keyword]; ok {
		return true
	}
	return false
}

func (L *LinkerMap) Add(keyword string) {
	L.Lock()
	L.internal[keyword] = struct{}{}
	L.Unlock()
	go func() {
		time.Sleep(L.timeout)
		L.Lock()
		delete(L.internal, keyword)
		L.Unlock()
	}()
}
