// Functions are first class types in Go.
// So a function can be assigned to a variable, passed to another function
// and returned from another function just like any other type
package main

import "fmt"

func main() {
	f := func(x int) {
		fmt.Println("Argument passed:", x)
	}
	// Here f is the function expression
	fmt.Printf("%T\n", f)
	// Invoke the function
	f(33)
}
