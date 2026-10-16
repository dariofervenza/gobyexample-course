package main

import (
	"fmt"
	"math"
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

}
