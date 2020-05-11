// dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	contents := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countDupLines(os.Stdin, contents)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countDupLines(f, contents)
			f.Close()
		}
	}

	for line, n := range contents {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countDupLines(f *os.File, contents map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		contents[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
