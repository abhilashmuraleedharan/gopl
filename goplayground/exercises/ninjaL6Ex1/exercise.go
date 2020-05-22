// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results
package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)
	y, s := bar()
	fmt.Println(y, s)
}

func foo() int {
	return 45
}

func bar() (int, string) {
	return 25, "Go is awesome"
}
