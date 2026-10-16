package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// if you have many arguments with the same type you can indicate it only in the last one
func plus2(a, b, c int) int {
	return a + b + c
}

// MULTIPLE return values

func dummy() (int, int) {
	return 1, 5
}

// variadic functions: can be called with any number or arguments (like fmt.Println(arg1, arg2, arg3, arg4, ...))

func variadic_fn(nums ...int) {
	total := 0
	fmt.Println("all nums are: ", nums)
	// nums is equivalent to []int
	for _, n := range nums {
		total += n
	}
	fmt.Println("Total is ", total)

}

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plus2(1, 2, 1))
	_, number := dummy()
	fmt.Println(number)

	variadic_fn(1, 1, 3)
	variadic_fn(1, 9)

	// we have the * equivalent from python to unpack a slice
	nuunms := []int{5, 6, 8}
	variadic_fn(nuunms...)
}
