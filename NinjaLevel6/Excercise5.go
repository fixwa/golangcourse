package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
}
type Circle struct {
	radius float64
}
type Shape interface {
	calculateArea() float64
}

func (s Square) calculateArea() float64 {
	return s.length * s.length
}

func (c Circle) calculateArea() float64 {
	return math.Pi * (c.radius * c.radius)
}

func info(s Shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	circ := Circle{radius: 12.34567}
	squa := Square{length: 22}

	info(circ)
	info(squa)
}
