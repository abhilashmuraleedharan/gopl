// To illustrate how to sort on fields in a struct using Sort func
// by implementing the sort package interface "Interface"
package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

// sort.Sort() function requires an Interface type value as its argument.
// I am defining a new type called byAge which is a []person
// that will implicitly implement the sort.Interface methods.
// When I do that byAge type also becomes Interface type.
type byAge []person

// Implement sort.Interface for byAge type
func (b byAge) Len() int {
	return len(b)
}
func (b byAge) Swap(i int, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byAge) Less(i int, j int) bool {
	return b[i].age < b[j].age
}

// Here I am creating a new type called byName
// to sort the []person by name
type byName []person

// Implement sort.Interface for byName type
func (b byName) Len() int {
	return len(b)
}
func (b byName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byName) Less(i, j int) bool {
	return b[i].name < b[j].name
}

func main() {
	p1 := person{
		name: "James Bond",
		age:  32,
	}

	p2 := person{
		name: "Miss Moneypenny",
		age:  27,
	}

	p3 := person{
		name: "Dr. Q",
		age:  37,
	}

	people := []person{p1, p2, p3}

	fmt.Println("[]people before sorting by age:")
	fmt.Println(people)
	sort.Sort(byAge(people))
	fmt.Println("[]people after sorting by age:")
	fmt.Println(people)
	fmt.Printf("\n")
	fmt.Println("[]people before sorting by name:")
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println("[]people after sorting by name:")
	fmt.Println(people)
}
