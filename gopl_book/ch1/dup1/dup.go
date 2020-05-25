// dup programs reads the lines entered in the standard input
// and prints the text of the lines that appear more than once preceded by their count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Make an empty map that takes string as key and int as value
	contents := make(map[string]int)
	// NewScanner function of bufio package produces a scanner type object
	// that reads given input and breaks it into lines or words
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // Each call to input.Scan() reads the next line and removes \n from end
		if len(input.Text()) > 0 { // Exit the program if an empty line is entered
			contents[input.Text()]++ // input.Text() retrieves the scanned value
		} else {
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()

	for line, n := range contents {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
