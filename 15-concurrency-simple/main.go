package main

import (
	"fmt"
	"sync"
)

func increment(counter *int, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		*counter++
	}
	wg.Done()
}

func main() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&counter, &wg)
	go increment(&counter, &wg)

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
