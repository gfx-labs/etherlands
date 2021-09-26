package main

import "time"

func start_repeating(duration time.Duration) chan bool {
	output := make(chan bool, 1)
	go (func() {
		for {
			time.Sleep(duration)
			select {
			case output <- true:
			default:
			}
		}
	})()
	return output
}

func start_delay(duration time.Duration) chan bool {
	output := make(chan bool)
	go (func() {
		time.Sleep(duration)
		output <- true
	})()
	return output
}
