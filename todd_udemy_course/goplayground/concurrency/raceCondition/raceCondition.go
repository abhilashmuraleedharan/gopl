// This program simulates the race condition depicted in the picture accessible from below link
// https://drive.google.com/file/d/0B22KXlqHz6ZNaW9mQ0U1b0tiUUE/view

// Race condition occurs when two parallely executing entities try to access the same shared memory
// The A race condition will give unpredictable results.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Number of Go routines:", runtime.NumGoroutine())

	const gr = 100 // Number of go routines to spin off

	counter := 0 // Shared variable

	var wg sync.WaitGroup
	wg.Add(gr) // Add the number of go routines to wait for

	// Simulate race condition
	for i := 0; i < gr; i++ {
		go func() {
			v := counter      // Read counter
			runtime.Gosched() // Yield thread
			v++               // Increment value
			counter = v       // Write counter
			wg.Done()         // Decrement the count of unfinished go routines
		}()
		fmt.Println("Number of Go routines:", runtime.NumGoroutine())
	}
	wg.Wait() // Wait for all go routines to finish
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
