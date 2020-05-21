// We can create a function which takes an unlimited number of arguments
// Parameter that can accept unlimited number of arguments is known as a “variadic parameter.”
// We use the lexical element operator ...T to specify variadic parameter of type T
package main

import "fmt"

func main() {
	total := sum(3, 5, 8, 9, 12, 45, 52)
	fmt.Println("Total: ", total)
}

// Variadic parameter should be the last parameter of a function
// Variadic parameter can also accept 0 arguments
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
