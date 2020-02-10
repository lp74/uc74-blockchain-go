package main

import (
	"fmt"
	"time"
)

func doStuff(quit chan bool, msg chan string, retry uint) {
	for retry > 0 {
		select {
		case <-quit:
			return
		default:
			retry--
			msg <- "."
			time.Sleep(time.Millisecond * 25)
		}
	}
}

func main() {
	quit := make(chan bool)
	msg := make(chan string)

	// Do stuff
	go doStuff(quit, msg, 100)

	// a timeout
	time.Sleep(time.Second)

	c := <-msg
	fmt.Println(c)

	// Quit goroutine
	quit <- true
}
