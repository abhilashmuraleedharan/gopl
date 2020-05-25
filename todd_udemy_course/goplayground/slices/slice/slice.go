// A SLICE holds VALUES of the same TYPE.
// A slice is created on top of an array
package main

import "fmt"

func main() {
	// We will ​use a COMPOSITE LITERAL to create a slice.
	x := []int{4, 5, 7, 8, 42}
	fmt.Println("x[1]: ", x[1])
	fmt.Println("Slice x: ", x)
	fmt.Println("x[1: ]", x[1:])
	fmt.Println("x[1:3]", x[1:3])

	// We can ​loop over the values in a slice with the range clause.
	for i, v := range x {
		fmt.Println(i, v)
	}
}
