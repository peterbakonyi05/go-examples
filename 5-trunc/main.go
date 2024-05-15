package main

import "fmt"

func main() {
	var a float32
	fmt.Println("Enter a floating point number:")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}
	result := int32(a)
	fmt.Println("Truncated version of the input:")
	fmt.Println(result)
}
