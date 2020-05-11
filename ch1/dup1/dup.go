// dup programs reads the lines entered in the standard input
// and prints the text of the lines that appear more than once preceded by their count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	contents := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		contents[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range contents {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
