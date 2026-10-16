package main

import "fmt"

func main() {
	// use range to sum numbers in a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum is ", sum)

	// range in arrays and slices returns index and the element
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index is: ", i)
		}
	}

	// range in maps returns key and value
	kvs := map[string]string{"a": "number 1", "b": "this is b"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate only the keys
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	// range on strings iterates over unicode points
	// first vlaue is the starting byte index of the rune
	// second is the rune itself

	fmt.Println("string range")
	for i, c := range "go a b" {
		fmt.Println(i, c) // 0 103  // 1 111
	}
}
