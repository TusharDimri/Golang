package main

/*
Mutex is Mutual Exclusion Lock.

It helps us lock the memory for one goroutines till it is working.

Read about Mutex and RW Mutex(Read Write Mutex) in Golang: https://pkg.go.dev/sync

Sometimes, there is a single DB or a varibale which is being used by multiple coroutines and this is where mutex comes into picture and saves us.
*/

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer

var signals = []string{"test"}
var mut sync.Mutex // pointer

func main() {

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

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}

}
