package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int, 5)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	// Receive only <-ch
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		wg.Done()
	}(myCh, wg)

	// Send only ch<-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
