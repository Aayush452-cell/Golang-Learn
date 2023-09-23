package main

import (
	"fmt"
	"time"
)

func main() {

	// unbuffered channel
	//ch := make(chan int)

	// buffered channel :- channel with a capacity, blocks only if the capaciity gets full.

	ch := make(chan int, 1)

	//fmt.Println("Sending value to channel")
	go send(ch)

	//fmt.Println("Recieving from channel")
	go receive(ch)

	time.Sleep(time.Second * 2)
}

func send(ch chan int) {
	ch <- 1
	//ch <- 2 no blockage till here in case of buffered as buffered channel we created had capcity 2
	//ch <- 3
	//close(ch)
	fmt.Println("Sended")
}

func receive(ch chan int) {
	time.Sleep(time.Second * 1)
	fmt.Println("Timeout finished")
	_, ok := <-ch
	fmt.Println(ok)
	//fmt.Printf("Value Recieved = %d", val)
}

// unbuffered channel :- channel can't store any data,doesn't have any storage.
// Timeout finished always before the Sended because send on an ubuffered channel is block until recieve happens on that channel in some other goroutine.
