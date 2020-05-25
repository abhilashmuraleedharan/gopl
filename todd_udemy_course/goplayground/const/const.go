package main

import "fmt"

// typed constants
const (
	a int     = 42
	b float64 = 42.67
	c string  = "Go Lang"
)

const hello = "Hello World" // untyped constant

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(hello)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", hello)
}
