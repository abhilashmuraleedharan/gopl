package main

import "fmt"

/*
 * Slices are built on top of arrays.
 * A slice is dynamic in that it will grow in size.
 * The underlying array, however, does not grow in size.
 * When we create a slice, ​we can use the built in function `make`
 * to specify how large our slice should be and also how large the underlying array should be.
 */
func main() {
	x := make([]int, 10, 12) // make([]T, length, capacity)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// When slice len becomes greater than size of the underlying array,
	// array size is doubled
	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
