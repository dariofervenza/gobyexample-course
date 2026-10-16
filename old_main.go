package main

import (
	"fmt"
	"math"
	"slices"
	"time"
)

func main() {
	fmt.Println("YOUR GAY")
	fmt.Println("VERY" + "GAY")
	fmt.Println("GAY", "YOU")
	fmt.Println("TRUE IS ", true)
	fmt.Println("TRUE && false IS ", true && false)
	fmt.Println("TRUE || false IS ", true || false)
	fmt.Println(7.6)

	// variables
	var a = 3

	var b, c int = 8, 7

	fmt.Println(a)
	fmt.Println(b, c)

	var d float32

	fmt.Println(d)

	d = 3.1

	fmt.Println(d)

	// use :=  ---> var x =
	f := "hello"
	fmt.Println("f is ", f)

	// constants

	const j = 5e20

	const div = 50000000

	const res = j / div

	fmt.Println(int64(res))
	fmt.Println(math.Acosh(res))

	// for loop represent the only iteration in go: for := for + while + loop

	i := 1

	// while equivalent
	for i <= 3 {
		fmt.Println("This is i:", i)
		i = i + 1
		i += 1
	}

	// classic for
	for d := 0; d < 5; d++ {
		fmt.Println("This is d:", d)
	}

	// loop equivalent

	for {
		break
	}

	// in range equivalent

	for r := range 6 {
		if r%2 == 0 {
			continue
		}
		fmt.Println("r is:", r)
	}

	// if else

	if 8%2 == 0 || 5-3 == 2 {
		fmt.Println("Gay")
	} else if 2-1 == -3 {
		fmt.Println("No gay")
	} else {
		fmt.Println("sUPER GAY")
	}

	// declare variable -> vlaid in the if else scope

	if if_var := 30; if_var < 20 {
		return
	} else if if_var < 40 {
		fmt.Println("Is less than 40")
	}

	// switch

	s_var := 50
	switch s_var {
	case 1:
		fmt.Println("switch 1")
	case 2, 50:
		fmt.Println("swithc 2 or 50")
	default:
		fmt.Println("switch rest")
	}

	switch time.Now().Weekday() {
	case time.Tuesday, time.Thursday:
		fmt.Println("Working at the office")
	default:
		fmt.Println("rest of the days")
	}

	// can be used as if, without indicating the variable
	t_n := time.Now()
	switch {
	case t_n.Hour() < 12:
		fmt.Println("Before noon")
	}

	// TYPE SWITCH
	whichIsMyType := func(val any) {
		// to get the type you could use also use reflect.TypeOf(val)
		switch t := val.(type) {
		case bool:
			fmt.Println("Im a boolean")
		case int:
			fmt.Println("im an integer")
		default:
			// notice the printf and also
			// t will hold the value not the type
			// if you want the type you need the %T
			// also without the \n printf does not add line jumps
			// NOTICE that %T with val would print the type
			fmt.Printf("I dont know the typeof: %T\n", t)
		}
	}
	whichIsMyType("hey")
	whichIsMyType(true)

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
