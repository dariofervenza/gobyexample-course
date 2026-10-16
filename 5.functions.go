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

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(plus2(1, 2, 1))
	_, number := dummy()
	fmt.Println(number)
}
