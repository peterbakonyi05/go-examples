package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Println("Hello from thread")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	// this will wait until there are 0 items in the wg
	wg.Wait()
	fmt.Println("Hello from main")
}
