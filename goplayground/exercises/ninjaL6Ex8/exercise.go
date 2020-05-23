// Create a func which returns a func
// Assign the returned func to a variable
// Call the returned func
package main

import "fmt"

func fun1() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	f := fun1()
	fmt.Printf("Type of f = %T\n", f)
	fmt.Println(f())
}
