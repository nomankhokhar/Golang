package main

import (
	"fmt"
	"time"
)

// pinger chan<- int, ponger <-chan int

// pinger chan<- int: This means that the pinger channel is a unidirectional channel that can only be used to send (chan<-) integer values. It can be used to send data to other goroutines, but you cannot receive data from it.

// ponger <-chan int: This means that the ponger channel is a unidirectional channel that can only be used to receive (<-chan) integer values. It can be used to receive data from other goroutines, but you cannot send data to it.

func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		value := <-pinger
		fmt.Println("ping", value)
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		value := <-ponger
		fmt.Println("pong", value)
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	select {}
}
