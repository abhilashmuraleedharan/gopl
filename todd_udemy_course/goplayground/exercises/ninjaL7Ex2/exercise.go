// create a person struct
// create a func called “changeMe” with a *person as a parameter
// change the value stored at the *person address
// create a value of type person & print out the value
// Call changeMe function
package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func changeMe(p *person) {
	// From Go Lang Specification
	// “As an exception, if the type of x is a named pointer type and (*x).f
	//  is a valid selector expression denoting a field (but not a method),
	//  x.f is shorthand for (*x).f.”
	p.firstname = "Richard"
	p.lastname = "Hammond"
}

func main() {
	p1 := person{
		firstname: "James",
		lastname:  "May",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
