package main

import (
	"fmt"
)

func send(c chan int) {
	c <- 3
	c <- 4
}

func main() {
	// create an int channel (second parameter is how many objects it can hold in transit, default is 0 which means not buffered)
	c := make(chan int, 1)
	go send(c)
	a := <-c
	b := <-c
	fmt.Println(a, b)
}
