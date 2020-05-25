// Go functions may be closures.
// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to these variables.
package main

import "fmt"

// This adder function is returning a closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x // referencing the variable sum
		// which is outside it's function body
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	// Here each closure will be bound to it's own sum variable.
	// Meaning the sum that pos bound to is in a different scope
	// from the sum neg is bound to which can be seen from the
	// below code output.
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
