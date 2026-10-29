package main

import (
	"fmt"
	"strconv"
)

// use strconv to parse strings into numbers
// works with (string_to_parse, bits_of_precision)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// this 0 means that it has to infer the base (base 2, base 10, hexadecimal)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// it auto recognizes hexadecimal
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// also parse as unsigned
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// USE ATOI TO TELL THAT THIS IS A BASE 10 int
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// if the string cannot be parsed it will return an err
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
