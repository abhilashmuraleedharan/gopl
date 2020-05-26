// Method sets determine what methods attach to a TYPE.
// It is exactly what the name says: What is the set of methods for a given type? That is its method set.

// IMPORTANT: “The method set of a type determines the INTERFACES
//             that the type implements.....”
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("Perimeter", s.perimeter())
}

func main() {
	c := circle{5}
	p := &c
	fmt.Println("Invoking area using value type c", c.area())
	fmt.Println("Invoking area using pointer type p", p.area())
	fmt.Println("Invoking perimeter using value type c", c.perimeter())
	fmt.Println("Invoking perimeter using pointer type p", p.perimeter())
	// Above statements show that a method set of a type can be accessed
	// using a pointer or non-pointer values of that type without any issue

	fmt.Println("Calling info with value type c as shape type")
	// info(c)  This does not work because area() requires a *T value since the receiver type is *T
	// when we try to invoke area() using an interface type (here shape)
	fmt.Println("Calling info with pointer type p as shape type")
	info(p) // This works because perimeter() can work with T and *T values
	// since the receiver type is T

	// Above statements show that when it comes to invoking methods of a type
	// using implemented interface type values, below rules matter!!
	// R1: Methods linked with receiver (t T) works with interface values of type T and *T
	// R2: Methods linked with receiver (t *T) works only with interface value of type *T
}
