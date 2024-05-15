package main

import "fmt"

// if you get it as a slice, it will not copy the whole array. More efficient.
// when size is specified, i.e.: [3]int, then this would copy the entire array.
func foo(sli []int) {
	sli[0] = sli[0] + 1
}

func main() {
	sli := []int{1, 2, 3}
	foo(sli)
	fmt.Println(sli)
}
