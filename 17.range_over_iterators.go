package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

// since go version 1.23

// they added support for iterators

// we reuse the List[T] form previous module

// equivaslent to AllElements but with iterator

// returns an iterator which is a func with a special signature (https://pkg.go.dev/iter#Seq)
func (list List[T]) All() iter.Seq[T] {
	// the iterator function takes as param another func
	// this func is called yield by convention
	// it will be called for every element we want to iterate over
	return func(yield func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !yield(e.val) {
				// early termination
				return
			}
		}
	}
}

// iteration does not require an underlying data structure and does NOT have to be FINITE
// here we iterate forever fibonacci series
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main_range_over_iterators(fib_max int) {
	list1 := List[int]{}
	list1.Push(4)
	list1.Push(40)
	list1.Push(89)
	for e := range list1.All() {
		fmt.Println(e)
	}

	// with package slices (https://pkg.go.dev/slices) we can conver any iterator into a slice
	// check slices, it has some more useful functions
	all := slices.Collect(list1.All())

	fmt.Println("All: ", all)

	// standard lib funtions also exposes iterators
	// eg strings.SplitSeq iterates over parts of a byte slice without declaring first a slice
	for part := range strings.SplitSeq("go-by-example", "-") {
		fmt.Printf("part: %s\n", part)
	}

	// infinite iterator usage example
	fmt.Println("\nFibonacci iterator:\n")
	for n := range genFib() {
		if n >= fib_max {
			break
		}
		fmt.Println(n)
	}

}
