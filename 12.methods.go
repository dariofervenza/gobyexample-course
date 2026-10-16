package main

import (
	"fmt"
)

// RUN go run 12.methods.go  13.interfaces.go  to compile both

// we can define methods on structs

type rectangle struct {
	width, height int
}

// metohds have to receive a pointer to the struct (not as param)
// its like it has 2 params field but first acts like &self in rust
func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perim() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := rectangle{3, 2}
	// now we can all them, they were assigned to rectangle

	fmt.Println("Area: ", r.area())
	fmt.Println("Perim: ", r.perim())

	rp := &r
	// even if a pointer, go manages conversion when calling methods (as with fields)
	fmt.Println("From a pointer area: ", rp.area())

	fmt.Println("\n\nINTERFACES:\n")
	main_interfaces()
}
