// Recursion happens when a function calls itself.
package main

import "fmt"

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	factorial := 1
	for ; n > 0; n-- {
		factorial *= n
	}
	return factorial
}

func main() {
	fmt.Println("4! =", factorial(4))
	fmt.Println("5! =", loopFactorial(5))
}
