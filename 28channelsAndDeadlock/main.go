//Channels are a way through which multiple go routines talk to each other

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)

	//Receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg) //Responsible for receiving the value


	//Send only 
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myCh <- 5
		// myCh <- 6
		myCh <-0
		close(myCh)
		close(myCh)
		wg.Done()
	}(myCh, wg) // Responsible for putting vales into the channel

	wg.Wait()
}
