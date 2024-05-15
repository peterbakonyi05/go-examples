package main

import (
	"fmt"
	"math"
)

func getFloatInput(prompt string) (float64, error) {
	var value float64
	fmt.Print(prompt)
	_, err := fmt.Scanf("%f", &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	a, err := getFloatInput("Enter acceleration: ")
	if err != nil {
		fmt.Println("Error reading acceleration:", err)
		return
	}

	v0, err := getFloatInput("Enter initial velocity: ")
	if err != nil {
		fmt.Println("Error reading initial velocity:", err)
		return
	}

	s0, err := getFloatInput("Enter initial displacement: ")
	if err != nil {
		fmt.Println("Error reading initial displacement:", err)
		return
	}

	fn := GenDisplaceFn(a, v0, s0)

	t, err := getFloatInput("Enter value for time: ")
	if err != nil {
		fmt.Println("Error reading time:", err)
		return
	}

	fmt.Println(fn(t))
}
