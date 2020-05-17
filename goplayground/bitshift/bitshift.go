package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1 // Left shift by 1 bit
	fmt.Printf("%d\t\t%b\n", y, y)
}
