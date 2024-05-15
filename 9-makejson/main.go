package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Println("Enter name:")
	fmt.Scanln(&name)
	fmt.Println("Enter address:")
	fmt.Scanln(&address)

	var data = map[string]string{
		"name":    name,
		"address": address}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}
