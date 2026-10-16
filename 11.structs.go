package main

import "fmt"

// structs are typed collections of fields
// useful for grouping data to form records

type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct

// go has a garbage collection so is safe to return a pointer to a local variable
// it will be cleaned from memory when there are no active references to it
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// also we can create it like this
	fmt.Println(person{"Bob", 20})

	// or used named fieldsa
	fmt.Println(person{name: "Thomas", age: 25})

	// omitted fields have zero value
	fmt.Println(person{name: "Fred"})

	// With & we get a pointer
	fmt.Println(&person{name: "Fred1"})

	// using the function, its idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Manolo"))

	// access fields with dot notation

	s := newPerson("Alice")
	fmt.Printf("\nName is %s and has age %d\n", s.name, s.age)

	// structs are mutable and you can modify them even if they are pointers (s is a pointer)
	s.age = 3
	fmt.Printf("\nAfter mods:\nName is %s and has age %d\n", s.name, s.age)

	// structs can also be anonymous
	// this is used in table-driven-tests: https://gobyexample.com/testing-and-benchmarking
	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println("Dog is: ", dog)

}
