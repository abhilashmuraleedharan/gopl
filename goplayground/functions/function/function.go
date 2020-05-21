package main

import "fmt"

func main() {
	fun1()                            // without argument
	fun2("James May")                 // with argument
	anchor := fun3("Richard Hammond") // with argument and a return value
	fmt.Println(anchor)
	anchor2, isFired := fun4("Jerremy Clarkson", "Top Gear") // with double arguments and double return values
	fmt.Println(anchor2)
	fmt.Println("IsFired: ", isFired)
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func fun1() {
	fmt.Println("Grand Tour")
}

// everything in Go is PASS BY VALUE
func fun2(s string) {
	fmt.Println("Hello ", s)
}

func fun3(s string) string {
	return fmt.Sprint("Hello ", s)
}

func fun4(anchor string, program string) (string, bool) {
	s := fmt.Sprint("Hello ", anchor, " from ", program)
	return s, true
}
