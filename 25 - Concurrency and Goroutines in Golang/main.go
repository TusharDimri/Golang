package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Concurrency vs Parallellism:

You have 3 tasks: A, B, C.

--> Concurrency:
A     A
   B      B
      C      C

--> Parallelism:
A  A  A  A
B  B  B  B
C  C  C  C

Goroutines are used to achieve parallelism in Golang.
   Threads               vs               Goroutines
Managed by OS                       Managed by Go Runtime
Fixed Stack - 1MB                   Flexible Stack - 2KB

" Do not communicate by sharing memory, instead share memory by communicating. "

*/

// Wait group
var wg sync.WaitGroup // pointers

func main() {
	// go greeter("Hello") // Goroutine
	// greeter("Good Morning")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://stackoverflow.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // Always at the end of the function
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // Tell that weight group that a task is done
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}

}
