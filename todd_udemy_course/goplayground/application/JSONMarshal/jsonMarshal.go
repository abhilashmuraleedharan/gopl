// Marshaling is the process of transforming the memory representation of an object to a data
// format suitable for storage or transmission, and it is typically used when data must be moved
// between different parts of a computer program or from one program to another.
package main

import (
	"encoding/json"
	"fmt"
)

// For all structs in GO, only fields with a capital first letter are
// visible to external programs/packages like the JSON Marshaller.
// Meaning that if you don't capitalize the first letter it won't get exposed
// when you convert it to JSON using the json.Marshall() function.

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "May",
		Age:   53,
	}

	p2 := person{
		First: "Jeremy",
		Last:  "Clarkson",
		Age:   59,
	}

	p3 := person{
		First: "Richard",
		Last:  "Hammond",
		Age:   45,
	}

	anchors := []person{p1, p2, p3}

	fmt.Printf("input type: %T\n", anchors)
	fmt.Println("input", anchors)

	// Using the Marshal function in json package, convert the []person to JSON
	// Marshal returns 2 values - the JSON string as a slice of bytes and error value if any
	// func Marshal(v interface{}) ([]byte, error)
	output, err := json.Marshal(anchors)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("output type: %T\n", output)
	fmt.Println("ouput after converting to string:", string(output))
}
