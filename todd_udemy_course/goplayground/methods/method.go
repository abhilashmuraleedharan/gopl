// A method is nothing more than a function attached to a TYPE.
// When you attach a function to a type, it is a method of that type.
// We atttach a function to a type with a RECEIVER

// Go function signature
// func (r receiver) identifier(parameters) (return(s)) { code }
package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstname, s.lastname)
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstname: "James",
			lastname:  "Bond",
		},
		licenseToKill: true,
	}

	sa2 := secretAgent{
		person: person{
			firstname: "Miss",
			lastname:  "Moneypenny",
		},
		licenseToKill: true,
	}

	sa1.speak()
	sa2.speak()
}
