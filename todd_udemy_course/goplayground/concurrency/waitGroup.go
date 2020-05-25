// In programming, concurrency is the ​composition​ of independently executing processes,
// while parallelism is the simultaneous ​execution​ of (possibly related) computations.

// Concurrency is about ​dealing with​ lots of things at once. Parallelism is about ​doing​ lots of things at once.
// Concurrent Programming is a design pattern where you write code that has the potential to be executed in parallel

// Go routine is similar to thread in Multi-Threaded programming in C/C++
// A WaitGroup waits for a collection of goroutines to finish​ just like join() waits in the main() in C/C++

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) // Sets the number of GO Routines to wait for
	go foo()  // Putting go in front of foo here will launch foo in its own go routine
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // Wait for all added GO routines to be finished
	fmt.Println("Exiting main")
}

func foo() {
	fmt.Println("Entering foo()")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Println("Exiting foo()")
	wg.Done() // Decrements the count of go routines yet to finish
}

func bar() {
	fmt.Println("Entering bar()")
	for i := 0; i < 10; i++ {
		fmt.Printf("%v\n", i)
	}
	fmt.Println("Exiting bar()")
}
