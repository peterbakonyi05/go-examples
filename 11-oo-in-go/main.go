// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Point struct {
// 	x float64
// 	y float64
// }

// // define the receiver type here (does not work for built in types)
// // by default it makes a copy
// func (p Point) DistanceFromOrigo() float64 {
// 	return math.Sqrt(p.x*p.x + p.y*p.y)
// }

// // note here: if you get the value without a reference, it will be copied!!
// // here because it changes point, it should be a pointer. No need to dereference here (automatic with the . operator)
// func (p *Point) Scale(v float64) {
// 	p.x = v * p.x
// 	p.y = v * p.y
// }

// func main() {
// 	// composite literal
// 	p := Point{x: 3, y: 4}
// 	p.Scale(2)
// 	fmt.Println("Distance from Origo: %d", p.DistanceFromOrigo())
// }
