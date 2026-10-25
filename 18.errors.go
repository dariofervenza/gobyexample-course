package main

import (
	"errors"
	"fmt"
)

// go communicates errors via an explicit separate return value

// go makes easy to know which functions return errors to handle them same way as other non error tasks

// see (https://pkg.go.dev/errors) and (https://go.dev/blog/go1.13-errors)

// by convention, errors are the last return value ahnd have type error (a built in interface)
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42")
	}
	// nil indicates that there was no error
	return arg + 3, nil
}

// a sentinel error is a predeclared variable that is used to express a specific error condition

var ErrOutOfTea = errors.New("run out of tea")
var ErrPower = errors.New("no power for tea")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// wrap the error with higher level errors
		// easiest if fmt.Errorf with %w
		// wrapped errors create a chain that can be queried wirg funcs like
		//		errors.Is
		//		errors.AsType
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// CUSTOM ERRORS: must implement Error() method

// use Error suffix
type argError struct {
	arg     int
	message string
}

// implementing the interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f132(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot run"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		// its idiomatic to check errors in a if
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// check if its a concrete error with errors.Is
			// will fit it the specified error exists in the chain
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("Buy tea please")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("No power")
			} else {
				fmt.Printf("unkown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is done")
	}

	fmt.Println("\n\nCUSTOM ERRORS:\n\n")

	_, err1 := f132(42)
	// errors.AsType is an advanced version of errors.Is
	// with checks if an error, or any in the chain matches an specific error type
	// and and converts to a error of that type, also retuning true
	// if there is no match, the second return is false
	if ae, ok := errors.AsType[*argError](err1); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("error is not argError")
	}

}
