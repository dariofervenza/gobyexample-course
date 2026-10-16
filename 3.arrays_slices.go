package main

import (
	"fmt"
	"slices"
)

func main() {

	// type assertion

	// only works if its an interface (uses any)
	var string_1 any = "gay"
	string_1_checked, ok := string_1.(string)

	if ok {
		fmt.Println(string_1_checked)
	}

	// arrays: sequence of elements with a specific length --> in go slices are more common in go

	var array_1 [5]int

	// it will be 5 zeros

	fmt.Println("Empty array:", array_1)

	// modify its values

	array_1[4] = 5
	array_1[0] = 1

	fmt.Println("Modified array:", array_1)

	// use len() as wit python
	fmt.Println(len(array_1))

	// to declare it with := you must follow this syntax
	array_2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_2)

	// or let the compiler count the elements

	array_3 := [...]int{6, 7, 8, 9, 10}
	fmt.Println(array_3)

	// you can assign one element to an index --> 3: 400 --> this means 400 goes in index = 3

	array_4 := [...]int{100, 3: 400, 500} // we define elements with index 0 , 3 and 4 --> hence elements 1 and 2 must be zero
	fmt.Println("Array 4 is: ", array_4)

	// 2D ARRAYS

	var two_array [2][3]int
	for idx1 := range 2 {
		for idx2 := range 3 {
			two_array[idx1][idx2] = idx1 + idx2
		}
	}
	fmt.Println("Original: ", two_array)

	//  reassign all values

	two_array = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(two_array)

	// slices: variable length arrays

	var slice_1 []string
	// if the slice is not initialited it will be equal to nil
	if slice_1 == nil {
		fmt.Println("Slice is empty with len =", len(slice_1))
	}

	// we can give a original capacity, but it can grow larger
	var slice_with_capacity = make([]string, 3)

	fmt.Println("Original values: ", slice_with_capacity)

	// make is a built in
	fmt.Println("Len: ", len(slice_with_capacity))
	fmt.Println("Capacity: ", cap(slice_with_capacity))

	slice_with_capacity[0] = "a"
	slice_with_capacity[1] = "b"
	slice_with_capacity[2] = "c"

	fmt.Println(slice_with_capacity)

	// append one value

	//notive that append returns a new slice
	slice_with_capacity = append(slice_with_capacity, "gay")
	slice_with_capacity = append(slice_with_capacity, "extra")

	// copy slices

	c8 := make([]string, len(slice_with_capacity))
	// destination, source
	copy(c8, slice_with_capacity)
	fmt.Println("copied is ", c8)

	// take a slice of a slice is similar to python

	l := c8[:3]
	l2 := c8[2:3]
	l3 := c8[3:]

	fmt.Println(l, l2, l3)

	// initialize with vlaues

	t12 := []int{1, 2}
	t13 := []int{1, 2}

	//slices.Equal() to check if they are equal

	if slices.Equal(t12, t13) {
		fmt.Println("They are equal")
	}

	// 2D slices

	two_d_slice := make([][]int, 3)
	for i2d := range 3 {
		inner_length := i2d + 1
		inner_slice := make([]int, inner_length)
		two_d_slice[i2d] = inner_slice
		for j2s := range inner_length {
			two_d_slice[i2d][j2s] = i2d + j2s
		}
	}
	fmt.Println("FINAL IS ", two_d_slice)

}
