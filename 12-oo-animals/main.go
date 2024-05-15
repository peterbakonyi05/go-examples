package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name             string
	food             string
	locomotionMethod string
	sound            string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotionMethod)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.sound)
}

func (a *Animal) Response(action string) string {
	switch action {
	case "eat":
		return a.food
	case "move":
		return a.locomotionMethod
	case "speak":
		return a.sound
	default:
		return "Unknown action"
	}
}

func main() {
	animals := map[string]Animal{
		"cow":   {name: "cow", food: "grass", locomotionMethod: "walk", sound: "moo"},
		"bird":  {name: "bird", food: "worms", locomotionMethod: "fly", sound: "peep"},
		"snake": {name: "snake", food: "mice", locomotionMethod: "slither", sound: "hss"}}

	// Reader to take input from the user
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter requests in the format '<animal> <action>' (e.g., 'cow speak')")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter in the format '<animal> <action>'.")
			continue
		}

		animalName, action := parts[0], parts[1]
		animal, ok := animals[animalName]
		if !ok {
			fmt.Printf("No data available for '%s'.\n", animalName)
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf("Unknown action: %s", action)
		}
	}
}
