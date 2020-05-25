// In Go, values can be of more than one type and an interface allows
// a value to be of more than one type.

// An interface can be empty(denoted as interface{}) or,
// it can have one or more function signatures within it.
// If a type has methods that matches these function signatures
// then we can say, that type implements this interface
// and value of that type is also a value of the interface type

// In other words, what an interface tells a type is,
// "If you have methods that match my functions, then you are my type"
package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak() // A type implementing interface human must have the method speak()
}

// Attaching speak func as a method of type person
func (p person) speak() {
	fmt.Println("I am", p.first, "and I am", p.age, "years old")
}

// Attaching speak function as a method of type secretAgent
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, "with license to kill attribute set to", s.ltk)
}

// Since both secretAgent and person types have method speak()
// we can say both the secretAgent and person types implement human interface
// and hence values of these types are also of type human.
func bar(h human) {
	// We  can switch on type as shown below
	switch h.(type) {
	case person:
		// Since now we know, this is a value of type person, we can assert it and access its field values
		fmt.Println("I am allowed inside the bar", h.(person).first, h.(person).age)
	case secretAgent:
		fmt.Println("I am allowed inside the bar", h.(secretAgent).first, h.(secretAgent).ltk)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			37,
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss Money",
			last:  "Penny",
			age:   27,
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "No",
		age:   40,
	}

	sa1.speak()
	sa2.speak()
	p1.speak()
	// Example of polymorphism
	bar(sa1)
	bar(sa2)
	bar(p1)
}
