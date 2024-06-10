// for value := range channel // this will run a for loop as long as the channel was not closed
// select to wait for one channel out of many to emit (might be send or receive as well)

package main

import (
	"fmt"
	"time"
)

// Worker function that sends data to a channel after some time
func worker1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Data from worker1"
}

// Another worker function that sends data to a channel after some time
func worker2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Data from worker2"
}

func main() {
	// Create two channels of type string
	ch1 := make(chan string)
	ch2 := make(chan string)
	abort := make(chan string)

	// Start two goroutines
	go worker1(ch1)
	go worker2(ch2)

	fmt.Println("Waiting for data from workers...")

	// Use select to wait on multiple channels
	// Note: it may be a mix or send/receive, whichever happens first will emit (i.e.: abort example here)
	select {
	case data := <-ch1:
		fmt.Println("Received:", data)
	case data := <-ch2:
		fmt.Println("Received:", data)
	case <-abort:
		// in case anything was sent to abort channel, quit
		return
		// this would run if none of the channels emit immediately
		// default:
		// 	fmt.Println("Noop")
	}
}
