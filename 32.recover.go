package main

import "fmt"

// it is possible to recover from a panic if we use recover() built-in

// recover can stop a program exit and let it continue

// example: a server does not fail even if a handler executes a panic
// in that case the server wnats to close the connection ands continue serving --> net/http does this out of the box

func mayPanic() {
	panic("scheduler panic")
}

func main() {
	// recover must be called in a deferred function
	// when the function panics it will activate the recover in the defer

	// the recover function returns the error that created the panic

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error: ", r)
		}
	}()

	mayPanic()

	// this will NOT RUN, after mayPanic the defer runs
	fmt.Println("After the panic")
}
