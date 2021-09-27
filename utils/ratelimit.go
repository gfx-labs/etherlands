package utils

import (
	"sync"
	"time"
)

type RateLimit struct {
	memory   map[uint64]time.Time
	duration time.Duration
	sync.RWMutex
}

func NewRateLimiter(duration time.Duration) *RateLimit {
	return &RateLimit{
		memory:   make(map[uint64]time.Time),
		duration: duration,
	}
}

func (R *RateLimit) Check(target uint64) bool {
	R.Lock()
	defer R.Unlock()
	if _, ok := R.memory[target]; ok {
		if time.Now().Sub(R.memory[target]) < R.duration {
			return false
		}
	}
	R.memory[target] = time.Now()
	return true
}
