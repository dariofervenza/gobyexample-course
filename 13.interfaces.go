package main

import (
	"fmt"
	"math"
)

// interfaces are collections of method signatures

// understand how interfaces work under the hood in https://research.swtch.com/interfaces

type geometry interface {
	area() float64
	// no comma needed
	perim() float64
}

// remember that go always uses the convention: name type  -> brand struct
type rectangle1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to comply with an interface we only need to implement all the methods in the interface

func (r rectangle1) area() float64 {
	return r.width * r.height
}

func (r rectangle1) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// implement an adapted using the geometry port interface
func measure(g geometry) {
	fmt.Printf("Data is type %T\nContent is %v\n", g, g)
	fmt.Println("Area: ", g.area(), "\nPerim: ", g.perim())
}

// add a different behavior depending on the implementation (circle or rectangle)
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Detected circle with radius: ", c.radius)
	}
}

func main_interfaces() {
	r := rectangle1{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
