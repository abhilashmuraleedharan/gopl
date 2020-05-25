// The sort package allows us to sort slices
package main

import (
	"fmt"
	"sort"
)

func main() {
	// define a slice of int to be sorted
	xi := []int{45, 67, 12, 5, 698, 23, 1, 99, 33}

	// define a slice of string to be sorted
	xs := []string{"Clarkson", "May", "Hammond"}

	fmt.Println("[]int before sorting:", xi)
	sort.Ints(xi)
	fmt.Println("[]int after sorting:", xi)

	fmt.Println("[]string before sorting:", xs)
	sort.Strings(xs)
	fmt.Println("[]string after sorting:", xs)
}
