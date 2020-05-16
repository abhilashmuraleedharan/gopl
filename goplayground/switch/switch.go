/*
 * In Go, we don't need to add break in every switch case as we do in C/C++
 * That's the default behavior. However, if you want few cases to fall through
 * use fallthrough keyword
 */
package main

import "fmt"

func main() {
	name := "Bond"
	switch name {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
		fallthrough
	case "R":
		fmt.Println("R?")
	default:
		fmt.Println("this is default")
	}
}
