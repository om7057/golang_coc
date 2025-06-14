package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in golang !")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3) //Indicating total number of go routines
	//Or can have 1 before each goroutine
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One Routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two Routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three Routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()


	fmt.Println(score)
}
