package main

/// lets do the same as we did in mutex but with channels

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// the state will be owned by only one goroutinem this guanrantees data is never corrupted
// the following structs encapsulate each of the read and write operations and a way for the owning goroutine to respond
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOpsCount uint64
	var writeOpsCount uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}

		}
	}()

	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				// track how many times we have read
				atomic.AddUint64(&readOpsCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOpsCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// sleep one second to let the functions run in backgorund, after this, man thread will end even if goroutines didnt finished (infinite loops)
	time.Sleep(time.Second)

	readOpsFInal := atomic.LoadUint64(&readOpsCount)
	writeOpsFInal := atomic.LoadUint64(&writeOpsCount)
	fmt.Println("READ OPS: ", readOpsFInal)
	fmt.Println("READ OPS: ", writeOpsFInal)

	// this stateful alternative is useful when you already have other channels or when you have to use multiple mutex
	// which is error-prone
}
