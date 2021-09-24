package main

import "time"



func start_repeating(milliseconds time.Duration) (chan bool) {
	output := make(chan bool, 1)
	go (func(){
		for{
			time.Sleep(milliseconds * time.Millisecond)
			select{
			case output <- true:
			default:
			}
		}
	})()
	return output
}



func start_delay(milliseconds time.Duration) (chan bool) {
	output := make(chan bool)
	go (func(){
		time.Sleep(milliseconds * time.Millisecond)
		output <- true
	})()
	return output
}
