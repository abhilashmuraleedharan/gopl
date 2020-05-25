// A struct is an composite data type.
// (composite data types, aka, aggregate data types, aka, complex data types).

// Structs allow us to compose together values of different types.
package main

import "fmt"

func main() {
	// Defining a new  type
	type person struct {
		firstname string
		lastname  string
		age       int
	}

	// Creating value of that type
	p1 := person{
		firstname: "James",
		lastname:  "Bond",
		age:       40,
	}
	fmt.Println(p1)
	fmt.Println(p1.firstname)

	// Embedding a type within another type
	// to create an embedded struct
	type secretAgent struct {
		person
		licenseToKill bool
	}

	sa1 := secretAgent{
		person: person{
			firstname: "James",
			lastname:  "Bond",
			age:       40,
		},
		licenseToKill: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.firstname, sa1.lastname, sa1.age, sa1.licenseToKill)
	// Here we can see inner type gets promoted to the outer type
	// meaning we don't need to specify sa.person.firstname here.
	// We need to be that explicit only if there is a name collision,
	// otherwise, sa.firstname would suffice.

	sa2 := secretAgent{
		person: person{
			firstname: "Jason",
			lastname:  "Bourne",
			age:       37,
		},
		licenseToKill: false,
	}
	fmt.Println(sa2)
	fmt.Println(sa2.firstname, sa2.lastname, sa2.age, sa2.licenseToKill)
}
