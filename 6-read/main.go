package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	var filePath string
	fmt.Println("File path with '[fname] [lname]' lines?")
	fmt.Scanln(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open the file")
		return
	}

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	names := make([]Person, 0, 0)

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		names = append(names, Person{fname: parts[0], lname: parts[1]})
	}

	for _, person := range names {
		fmt.Printf("Firstname: %s, Lastname: %s\n", person.fname, person.lname)
	}

	file.Close()
}
