package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// Use append function to delete elements from a slice as well
	// Here I am removing the elements 77, 88, 99 and 1014 from slice
	x = append(x[:5], x[9:]...)
	fmt.Println(x)
}
