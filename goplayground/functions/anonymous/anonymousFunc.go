// Anonymous self executing functions
package main

import "fmt"

func main() {
	// Anonymous function that takes no arguments
	func() {
		fmt.Println("Anonymous func with no arguments ran")
	}() // the parenthesis here invokes the anonymous function

	// Anonymous function that takes an argument
	func(x int) {
		fmt.Println("Anonymous function ran with int argument:", x)
	}(13)
}
