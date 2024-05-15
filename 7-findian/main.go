package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter a string")
	fmt.Scanln(&input)
	if len(input) > 2 && input[0] == 'i' && strings.Contains(input, "a") && input[len(input)-1] == 'n' {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
