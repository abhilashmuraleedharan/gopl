// When we have a slice of some type, we can pass in the individual values in a slice by
// using the “...” operator.
package main

import "fmt"

func main() {
	xi := []int{3, 5, 8, 9, 12, 45, 52}
	total := sum(xi...)
	fmt.Println("Total: ", total)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
