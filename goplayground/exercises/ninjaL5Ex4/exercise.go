// Write a program to demonstrate anonymous struct
package main

import "fmt"

func main() {
	// Anonymous struct
	s := struct {
		name      string
		friends   map[string]string
		favDrinks []string
	}{
		name: "Richard Hammond",
		friends: map[string]string{
			"Jerremy Clarkson": "5678812345",
			"James May":        "2567891236",
		},
		favDrinks: []string{
			"Gin and Tonic",
			"Water",
		},
	}
	fmt.Println(s.name)
	fmt.Println("Friends: ")
	for k, v := range s.friends {
		fmt.Println(k, ":", v)
	}
	fmt.Println("Favorite drinks: ")
	for _, drink := range s.favDrinks {
		fmt.Printf("%s, ", drink)
	}
}
