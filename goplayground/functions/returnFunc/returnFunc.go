// Functions are first class types in Go and hence
// can be returned from a function just like any other type
package main

import "fmt"

func main() {
	fmt.Println(foo()()) // Here the returned function is immediately invoked within the Println
}

// Here foo() function returns a function
// that takes no argument and returns an int value
func foo() func() int {
	return func() int {
		return 1729 // Ramanujan's number
	}
}
