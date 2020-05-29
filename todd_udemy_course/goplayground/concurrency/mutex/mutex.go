// A “mutex” is a mutual exclusion lock.
// Mutexes allow us to lock our code so that only one goroutine can access that locked chunk of code at a time.
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

	var mu sync.Mutex

	// Simulate race condition
	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter      // Read counter
			runtime.Gosched() // Yield thread
			v++               // Increment value
			counter = v       // Write counter
			mu.Unlock()
			wg.Done() // Decrement the count of unfinished go routines
		}()
		fmt.Println("Number of Go routines:", runtime.NumGoroutine())
	}
	wg.Wait() // Wait for all go routines to finish
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
