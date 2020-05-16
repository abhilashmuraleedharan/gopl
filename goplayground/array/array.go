// Main thing to note about arrays in Go is size is part of the type as illustrated below
package main

import "fmt"

func main() {
	var array1 [25]int
	array1[5] = 37
	var array2 [15]int
	array2[14] = 99
	fmt.Println("array1", array1, "len", len(array1))
	fmt.Printf("array1 type: %T\n", array1)
	fmt.Println("array2", array2, "len", len(array2))
	fmt.Printf("array2 type: %T\n", array2)
}

/*
 * Arrays are mostly used as a building block in the Go programming language.
 * In some instances, we might use an array,
 * but mostly the recommendation is to ​use slices instead.
 */
