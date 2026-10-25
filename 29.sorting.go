package main

import (
	"cmp"
	"fmt"
	"slices"
)

// which types are ordered? check  cmp.Ordered (https://pkg.go.dev/cmp#Ordered)
func main() {
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{1, 59, 74, 3, 5965, 8}
	slices.Sort(ints)
	fmt.Println("Insts: ", ints)

	s := slices.IsSorted(strs)
	fmt.Println("Is sorted? ", s)

	// SORTING BY KEY (OR FUNCTION)

	// eg. sort a []string by each element length

	fruits := []string{"peach", "banana", "hi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)

	fmt.Println("Sorted fruits: ", fruits)

	// sort slice of structs

	type Person12 struct {
		name string
		age  int
	}

	people := []Person12{
		Person12{name: "Jax", age: 37},
		Person12{name: "Thomas", age: 25},
		Person12{name: "Alexander", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person12) int {
			return cmp.Compare(a.age, b.age)
		},
	)
}
