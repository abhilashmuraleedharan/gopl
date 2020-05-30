// In addition to the main goroutine, launch two additional goroutines
//     * each additional goroutine should print something out
//     * use waitgroups to make sure each goroutine finishes before your program exists
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Starting in the main goroutine")
	// define sync waitgroup variable to synchronize Go routines execution
	var wg sync.WaitGroup

	// Launch 2 go routines
	wg.Add(2) // Number of Go routines to be present in the wait group
	for i := 0; i < 2; i++ {
		go func(i int) {
			fmt.Println("Launched new Go routine", i+1)
			runtime.Gosched()
			wg.Done()
		}(i)
	}
	// Wait for the go routines to finish executing
	wg.Wait()
	fmt.Println("Exiting main go routine")
}
