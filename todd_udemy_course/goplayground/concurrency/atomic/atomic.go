// Atomic package contain methods that we can use to perform atomic operations on a shared memory variable.
// This is another way of preventing race condition without resorting to mutex.
// As per the Go Lang Spec atomic package methods should only be use for some special low level programming
// and should be used with great care. For most concurrent programming synchronization can be achieved through channels.

// Remember the Go Lang's Motto on Concurrency:
// "Share memory by communicating; Don't communicate by sharing memory."
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Number of Go routines:", runtime.NumGoroutine())

	const gr = 100 // Number of go routines to spin off

	var counter int64 // To use the atomic package methods AddInt64() and LoadInt64()

	var wg sync.WaitGroup
	wg.Add(gr) // Add the number of go routines to wait for

	// Simulate race condition
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched() // Yield thread
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done() // Decrement the count of unfinished go routines
		}()
		fmt.Println("Number of Go routines:", runtime.NumGoroutine())
	}
	wg.Wait() // Wait for all go routines to finish
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
