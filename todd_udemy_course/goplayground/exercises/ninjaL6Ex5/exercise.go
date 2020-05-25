// create a type ​SQUARE
// create a type ​CIRCLE
// attach a method to each that calculates ​AREA​ and returns it
// create a type ​SHAPE​ that defines an interface as anything that has the ​AREA​ method
// create a func ​INFO​ which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle
package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("Area of square:", s.area())
	case circle:
		fmt.Println("Area of circle:", s.area())
	}
}

func main() {
	s := square{
		side: 4.7,
	}

	c := circle{
		radius: 6.3,
	}

	info(s)
	info(c)
}
