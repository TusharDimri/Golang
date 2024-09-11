package main

/*
A RWMutex provides more granularity:
It allows multiple goroutines to read a resource at the same time (shared reads).
But only one goroutine can write at any given time (exclusive write).
RWMutex provides two sets of locks:
Read lock: RLock() - Multiple readers can hold this lock simultaneously.
Write lock: Lock() - Only one writer can hold this lock at a time, and no readers are allowed during a write.
When to use:
Use RWMutex when you have a scenario where many reads but few writes happen. It allows multiple goroutines to read simultaneously without blocking each other, but ensures that write operations are exclusive.
*/

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(6) // We already know the number of coroutines

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 4")
		mut.Lock()
		score = append(score, 4)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 5")
		mut.Lock()
		score = append(score, 5)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	//  Immediately executing function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		fmt.Println("Goroutine 6")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
	//  go run --race ."file name" ( It will help you test for Race Condition )
}
