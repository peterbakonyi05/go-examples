package main

import "fmt"

func main() {
	// go keyword starts a new go subrouting managed by the go runtime scheduler
	// like this it will not print the thread because main finishes first
	// when main thread finishes, all the subroutines are killed
	go fmt.Println("Hello from thread")
	fmt.Println("Hello from main")
}
