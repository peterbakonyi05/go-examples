package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

// no need to state explicitly that it satisfies the interface, compiler will do the matching
type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) Area() float64 {
	s := float64(t.a+t.b+t.c) / 2.0
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func main() {
	t := Triangle{a: 3, b: 5, c: 4}
	fmt.Println(t.Area())

	var s Shape2D
	s = t
	fmt.Println(s.Perimeter())
}
