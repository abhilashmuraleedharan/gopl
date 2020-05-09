// This program illustrates a different variant of echo program
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", "" // short variable declaration form

	/*
	 * Range based for loop
	 * Use the blank identifier `_` whenever syntax requires
	 * a variable name but the program logic does not.
	 * Here, range produces a pair of values, index and the value at that index
	 * Since we don't need the index in our logic, we can use the blank identifier
	 */
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
