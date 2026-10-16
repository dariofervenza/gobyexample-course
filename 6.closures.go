package main

import "fmt"

// closures are anonymous functions, used to define a function without a name

// we define a function that returns another funtion

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// uses a different i value
	nextIntNew := intSeq()
	fmt.Println(nextIntNew())
}
