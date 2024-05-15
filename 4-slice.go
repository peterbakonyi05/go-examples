// package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	slice := make([]int, 0, 3)
// 	var input string

// 	for {
// 		fmt.Scanln(&input)
// 		if strings.ToUpper(input) == "X" {
// 			break
// 		}
// 		num, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("Error converting string to int:", err)
// 			continue
// 		}
// 		slice = append(slice, num)
// 		sort.Ints(slice)
// 		fmt.Println(slice)

// 	}
// 	fmt.Println("Finished")
// }
