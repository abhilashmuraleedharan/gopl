/*
 * Within a constant declaration, the predeclared identifier iota represents
 * successive untyped integer constants.
 * It is reset to 0 whenever the reserved word const appears in the source.
 * It can be used to construct a set of related constants
 */
package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)

const (
	g = iota
	h = iota
	i = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}
