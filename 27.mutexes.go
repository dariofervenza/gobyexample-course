package main

import (
	"fmt"
	"sync"
)

// for operations more complex than increasing a atomic counter we can use a mutex to safely access data across goroutines

// hold a map of counters
// we can synchronize access usign the mutex
// NOTICE THAT: mutexes are not for copying, if we pass the struct we must do it with a pointer
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// lock the mutex
	c.mu.Lock()
	// unlock it at the end of the func with defer statement (https://gobyexample.com/defer)
	// defer is like finally in python (defer closeFile(f) === finally: file.close())
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// we can use the zero value of a mutex, no explicit declaration
	cMu := Container{
		counters: map[string]int{"a": 0, "b": 1},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			cMu.inc(name)
		}
	}

	wg.Go(func() {
		doIncrement("a", 100000)
	})
	wg.Go(func() {
		doIncrement("b", 400000)
	})
	wg.Wait()
	fmt.Println(cMu.counters)
}
