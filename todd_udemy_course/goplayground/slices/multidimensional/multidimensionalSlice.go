package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	z := [][]int{x, y}
	fmt.Println("slice x: ", x)
	fmt.Println("slice y: ", y)
	fmt.Println("slice z: ", z)
	fmt.Printf("z type: %T\n", z)
}
