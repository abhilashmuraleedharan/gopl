// The inverse of marshaling is called unmarshaling or demarshaling.

package main

import (
	"encoding/json"
	"fmt"
)

// I am using the site https://mholt.github.io/json-to-go/ to determine the Go type
// for a given JSON. I got the below output for the JSON {"First":"James","Last":"May","Age":53}
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// In Golang when you put a string in backtick ``, it is treated as raw string literal.
	// Meaning that it is read exactly as is, a UTF-8 string with no escape characters
	// only runes(utf-8 code points) Vs the "" double quotes which allows for escaped characters.
	ip := `[{"First":"James","Last":"May","Age":53},{"First":"Jeremy","Last":"Clarkson","Age":59},{"First":"Richard","Last":"Hammond","Age":45}]`
	fmt.Println("input as a string:", ip)

	// When we want to convert a JSON Object into a Golang struct, we use the json.Unmarshal.
	// Umarshal is Golangs way of saying "parse this JSON object into a valid GOlang struct".
	// Unmarshal takes the following parameters:
	// 1. a slice of bytes (This a raw string, this is the JSON object that we want to parse)
	// 2. A pointer to a struct to parse the JSON into

	// Prepare argument 1: []bytes containing raw JSON string to parse
	input := []byte(ip)
	fmt.Println("input as a slice of bytes:", input)

	// Prepare argument 2: Struct to parse the JSON into
	var anchors []person

	// func Unmarshal(data []byte, v interface{}) error
	err := json.Unmarshal(input, &anchors)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Output type: %T\n", anchors)
	fmt.Println("output:")

	for i, v := range anchors {
		fmt.Printf("%v. %v %v of age %v\n", i+1, v.First, v.Last, v.Age)
	}
}

// ONLY FIELDS FOUNDS IN THE destination type(struct) will be decoded.
// Meaning if there is a field in the JSON that isn't in the destination it will be ignored.
