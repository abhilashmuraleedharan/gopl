package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	// To append values to a slice, we use the special built in function append
	// This function returns a slice of the same type.
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...) // The ... operator if used after variable name means unfurl it
	fmt.Println(x)
}
