package main

import "time"





type ratelimit struct {
	memory map[uint64]time.Time
}


func NewRateLimiter() *ratelimit {
	return &ratelimit{
		memory: make(map[uint64]time.Time),
	}
}


func (R *ratelimit) check(target uint64) bool {
	if _, ok := R.memory[target]; ok {
		if time.Now().Sub(R.memory[target]).Seconds() < 30 {
			return false;
		}
	}
	R.memory[target]  = time.Now();
	return true;
}
