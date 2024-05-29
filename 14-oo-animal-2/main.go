package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		commands := strings.Fields(input)

		if len(commands) != 3 {
			fmt.Println("Invalid command")
			continue
		}

		switch commands[0] {
		case "newanimal":
			name := commands[1]
			animalType := commands[2]

			var animal Animal
			switch animalType {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Invalid animal type")
				continue
			}

			animals[name] = animal
			fmt.Println("Created it!")
		case "query":
			name := commands[1]
			action := commands[2]

			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found")
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
				fmt.Println("Invalid query action")
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
