package main

import "fmt"

// use pointers to pass references to values and records in your programm

// zeroval accepts ival as value
// so ival inside zeroval its a copy of the original variable
func zeroval(ival int) {
	ival = 0
}

// zeroptr has *int so it takes an int pointer
// *iptr dereferences the pointer from memory address to the current value of the address
// so modifying it modifies the original value
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Original value is ", i)
	zeroval(i)
	fmt.Println("Zeroval: ", i)

	// with &i we send the memory address, this is a pointer
	zeroptr(&i)
	fmt.Println("Zeroptr: ", i)
	fmt.Println("Pointer: ", &i)

	// a pointer can also be created with new()

	p := new(42)
	fmt.Println("value at *p: ", *p)
	zeroptr(p)
	fmt.Println("value after zero prt at *p: ", *p)

}
