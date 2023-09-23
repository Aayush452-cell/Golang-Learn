package main

import (
	"fmt"
	"time"
)

func execute(id int) {
	fmt.Println("GoRoutine ", id)
}

func temp5() {
	/* go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished") */

	// Mutiple go routines to show go routines are indepented and exceute concurrently
	/* fmt.Println("Started")
	for i := 0; i < 10; i++ {
		go execute(i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("Finished") */

	//fmt.Println(runtime.NumCPU())

	// Anonymous Go routine

	go func() {
		fmt.Println("In Go routine")
	}()

	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {
	go start2()
	fmt.Println("In Goroutine 1")
}

func start2() {
	fmt.Println("In Goroutine 2")
}
