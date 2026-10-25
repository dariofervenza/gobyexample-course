package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// share data between goroutines using the sync/atomic package

func main() {
	// create an atomic integer to represent our positive counter
	var ops atomic.Uint64

	var wg sync.WaitGroup

	// IONCREATE THE COUNTER 1000 times each of the 50 workers
	for range 50 {
		wg.Go(func() {
			for range 1000 {
				ops.Add(1)
			}
		})
	}

	// wait them to finish
	wg.Wait()

	// it should be 50 000
	// load is safe to read even if other goroutines are atomically updating it
	fmt.Println("Ops final result: ", ops.Load())

	// if we didnt used an atomic counter, we would get a different number than 50000
	// if we run without atomic with -race flag (go run module.go -race) wwe would get data race failures
}
