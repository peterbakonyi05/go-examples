// interleaving happens at a machine code level! Not at go source code level. Makes sense.
package main

import (
	"fmt"
	"sync"
)

var i = 0
var m sync.Mutex
var wg sync.WaitGroup

func inc() {
	m.Lock()
	i = i + 1
	m.Unlock()
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
