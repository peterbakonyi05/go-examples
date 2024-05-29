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

type Rectangle struct {
	a float64
	b float64
}

func (t Rectangle) Area() float64 {
	return t.a * t.b
}

func (t Rectangle) Perimeter() float64 {
	return 2 * (t.a + t.b)
}

func DrawShape(s Shape2D) {
	// type assertion with if else
	tri, ok := s.(Triangle)
	if ok {
		fmt.Println("Triangle", tri)
	} else {
		fmt.Println("Not a triangle", s)
	}
}

// type switch
func DrawShape2(s Shape2D) {
	switch sh := s.(type) {
	case Rectangle:
		fmt.Println("RECT", sh)
	case Triangle:
		fmt.Println("TRI", sh)
	}
}

// interface{} is the any type. There is also an `any` alias type
func DrawAnything(s interface{}) {
	switch sh := s.(type) {
	case Rectangle:
		fmt.Println("RECT", sh)
	case Triangle:
		fmt.Println("TRI", sh)
	}
}

func main() {
	t := Triangle{a: 3, b: 5, c: 4}
	fmt.Println(t.Area())

	var s Shape2D
	s = t
	fmt.Println(s.Perimeter())
	DrawShape(s)

	/** Dynamic type and value */
	var s2 Shape2D
	DrawShape(s2)
	// this will fail because the dynamic type is null
	// fmt.Println("With dynamic type not set, this will throw a runtime error",s2.Area())
	var rect Rectangle
	s2 = rect
	// at this point the dynamic type is set only the dynamic value is null so it does not fail
	fmt.Println("With dynamic type set, this will return the default value", s2.Area())
}
