// Share memory by communicating don't communicate by sharing memory

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup //Usually pointers
var mut sync.Mutex //Pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://go.dev",
		"https://pkg.go.dev",
		"https://google.com",
		"https://microsoft.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for %s\n ", result.StatusCode, endpoint)
	}
}
