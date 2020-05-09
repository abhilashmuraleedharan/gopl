// echo program using String Join function
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // More efficient than the previous variants
	// fmt.Println(os.Args[1:])                    // Any slice may be printed this way
}
