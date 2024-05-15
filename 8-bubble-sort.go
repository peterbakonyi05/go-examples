// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func Swap(slice []int, i int) {
// 	if i < len(slice)-1 {
// 		slice[i], slice[i+1] = slice[i+1], slice[i]
// 	}
// }

// func BubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j] > slice[j+1] {
// 				Swap(slice, j)
// 			}
// 		}
// 	}
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter up to 10 integers, separated by spaces:")
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}
// 	inputs := strings.Fields(input)
// 	ints := make([]int, len(inputs))
// 	for i, str := range inputs {
// 		num, err := strconv.Atoi(str)
// 		if err != nil {
// 			fmt.Println("Error parsing integer:", err)
// 			return
// 		}
// 		ints[i] = num
// 	}
// 	BubbleSort(ints)

// 	fmt.Println("Sorted integers:")
// 	for _, num := range ints {
// 		fmt.Print(num, " ")
// 	}
// 	fmt.Println()

// }
