// Call function is a function that's passed into another function
// as an argument. Functional Programming is not something recommended in Go.
// Because idiomatic go means clear, simple, readable code.
// Below is just an example of how we can do functional programming using Go.
package main

import "fmt"

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Example of functional programming
// This function will return the sum of only even numbers
func evenSum(f func(...int) int, xi ...int) int {
	// Extract even numbers from input slice
	var ex []int

	for _, v := range xi {
		if v&1 == 0 { // Check if even
			ex = append(ex, v)
		}
	}
	// ex slice will now have only even numbers.
	// Unfurl ex slice and pass into f to get the sum
	return f(ex...)
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Input slice: ", ii)
	fmt.Println("Sum of all numbers: ", sum(ii...))
	fmt.Println("Sum of only even numbers: ", evenSum(sum, ii...))
}
