// echo program prints on the console its command line arguments
// like the Unix echo command
package main

import (
	"fmt"
	"os"
)

/*
 * Command-line arguments are available to a program in a variable named Args that is part of the os package
 * The variable os.Args is like a slice of strings.
 * The first element of os.Args, os.Args[0] is the name of the command itself, the other elements are the
 * arguments presented to the program
 */
func main() {
	var s, sep string // here sep will act as separator

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
