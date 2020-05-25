// A pointer is a variable that holds the address of a memory location
package main

import "fmt"

func main() {
	var x int
	fmt.Printf("x = %v\n", x)
	fmt.Printf("Address of x = %v\n", &x)
	p := &x
	fmt.Println("Pointer p value = ", p)
	fmt.Printf("Type of p = %T\n", p)
	var q *int
	fmt.Println("Make pointer q point to same memory location pointed by p")
	q = p
	fmt.Println("Pointer q value = ", q)
	fmt.Printf("Type of q = %T\n", q)
	fmt.Println("Storing 10 in the memory address pointed by q")
	*q = 10
	fmt.Println("Value of x now:", x)
	fmt.Println("Calling fun1 by passing x")
	fun1(x)
	fmt.Println("Value of x after calling fun1:", x)
	fmt.Println("Calling fun2 by passing the address of x by value")
	fun2(&x)
	fmt.Println("Value of x after calling fun2:", x)
}

// Non mutating function
func fun1(x int) {
	x++
	fmt.Println("Printing value of x from fun1:", x)
}

// Mutating function
func fun2(xp *int) {
	*xp++
	fmt.Println("Printing value stored in xp from fun2:", *xp)
}

// In Go, everything is Pass by Value.
// There is no copy argument by value, copy argument by reference etc. in Go
// In above example, we are passing the address of x by value
