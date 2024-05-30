package main

import (
	"fmt"
	"sync"
)

func prod(v1 int, v2 int, c chan int, wg *sync.WaitGroup) {
	// this writes into the channel (blocking instruction until the receive happens! by default there is no buffering)
	// note: this is also a way of synchronisation because this will wait until continues
	c <- v1 * v2
	// this will only be printed af the channel has been read and only if there is a WaitGroup
	// otherwise the main thread completes and this never runs (pretty complex...)
	fmt.Println("Finished prod", v1, v2)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// create an int channel (second parameter is how many objects it can hold in transit, default is 0 which means not buffered)
	c := make(chan int)
	wg.Add(1)
	go prod(1, 2, c, &wg)
	wg.Add(1)
	go prod(3, 4, c, &wg)
	fmt.Println("This comes before print within prod because thread is waiting until someone is reading the value from the channel")
	a := <-c // most recent
	b := <-c // second most recent
	fmt.Println(a * b)
	wg.Wait()
}
