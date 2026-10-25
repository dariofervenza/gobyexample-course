package main

import (
	"fmt"
	"time"
)

// a goroutine is a thread of execution

func fGoRou(from string) {
	for i := range 8 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// synchonous execution
	fGoRou("direct")

	// invoke in a goroutine with go functName()

	go fGoRou("goroutine")

	// start goroutine from anonymous func

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// for a more robust approach use a WaitGroup (https://gobyexample.com/waitgroups)
	time.Sleep(time.Second)
	fmt.Println("done")
}
