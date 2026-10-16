package main

import "fmt"

// RUN WITH: go run 16.generics.go  17.range_over_iterators.go

// since go 1.18

// generics are also nkow as type parameters

// example of a generic function
// takes a slice of any comparable type
// and a element of that type

// the comparable constraint means that we can compare values of it with the operators:
// 		==
// 		!=

// see https://go.dev/blog/deconstructing-type-parameters
// and notice htat this function (same implementation) exists as slices.Index (https://pkg.go.dev/slices#Index)

func SlicesIndex[SliceType ~[]ElementType, ElementType comparable](s SliceType, v ElementType) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// example of a generic in a linked list

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// we can use them in functions remember that the type is not List --> it is List[T]
// add a method to the struct
func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

// get all elements of the linked list
func (list *List[T]) AllElements() []T {
	var elements []T
	// NOTICE THIS FOR LOOP --> e = e.next
	for e := list.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func main() {
	// generic functions invocations often do type inference, we do not need to specify them
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo: ", SlicesIndex(s, "zoo"))

	// we can pass the generic types explicitly
	_ = SlicesIndex[[]string, string](s, "bar")

	// linked list usage

	list := List[int]{}
	list.Push(10)
	list.Push(13)
	list.Push(23)

	fmt.Println("All elements:\n", list.AllElements())

	fmt.Println("\n\nRange Over Iterators\n\n")
	main_range_over_iterators(50)

}
