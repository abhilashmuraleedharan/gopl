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
	if len(files) == 0 { // Revert to reading lines from stdin if no file names are passed
		countDupLines(os.Stdin, contents)
	} else {
		/*
		 * In each iteration read the file name from Args,
		 * open the file and read the lines in it
		 * close the file if file was opened successfully
		 */
		for _, arg := range files {
			/*
			 * os.Open returns 2 values.
			 * First one is the open file (*os.File)
			 * Second one is the result of os.Open which is a value of built-in error type
			 */
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countDupLines(f, contents)
			f.Close()
		}
		// NOTE: ignoring potential errors from os.Open() to keep the e.g simple
	}

	// Print the duplicate lines read with the count
	for line, n := range contents {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Read lines from the given input and update the map datastructure
func countDupLines(f *os.File, contents map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if len(input.Text()) > 0 {
			contents[input.Text()]++
		} else {
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
