// The previous 2 versions of dup operate in a streaming mode in which input is read
// and broken into lines as needed, so in principle those programs can handle an arbitrary
// number of input.
// In this program an alternate approach is taken wherein the program reads the entire input
// into memory in one big gulp, split into lines  all at once, then process the lines.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	contents := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile function reads the entire contents of a named file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// ReadFile returns a byte slice so it must be converted into a string
		// so that it can be split by strings.Split()
		for _, line := range strings.Split(string(data), "\n") {
			if len(line) > 0 {
				contents[line]++
			}
		}
	}
	for line, n := range contents {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
