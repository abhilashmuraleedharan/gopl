package main

import "fmt"

func main() {
	// Defining a new type
	type person struct {
		firstname               string
		lastname                string
		favoriteIcecreamFlavors []string
	}

	// Creating values of that type
	p1 := person{
		firstname:               "Jerremy",
		lastname:                "Clarkson",
		favoriteIcecreamFlavors: []string{"Pista", "Strawberry", "Butterscotch"},
	}

	p2 := person{
		firstname:               "James",
		lastname:                "May",
		favoriteIcecreamFlavors: []string{"Coconut", "Vanilla", "Belgium Chocolate"},
	}

	// Creating a map to store values of the new type
	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	// Printing the contents of the map
	for k, v := range m {
		fmt.Println("For key: ", k)
		fmt.Println("First Name: ", v.firstname)
		fmt.Println("Last Name: ", v.lastname)
		fmt.Printf("Fav Icecream Flavours: ")
		for _, f := range v.favoriteIcecreamFlavors {
			fmt.Printf("%s, ", f)
		}
		fmt.Printf("\n")
	}
}
