package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("part 2")
	// maps
	mymap := make(map[string]int) // [key type]inner type
	// update map
	mymap["gay"] = 0
	mymap["your"] = 1

	fmt.Println("Map: ", mymap)

	// get a value

	value := mymap["gay"]
	fmt.Println("value is", value)

	// if a key does not exist, i get the ZERO VALUE
	// ZERO VALUE is false, 0, "" or nil for ppointers functions interfaces, slices, channels, and maps

	value_not_exists := mymap["idontexist"]
	fmt.Println("value_not_exists is", value_not_exists)

	// number of elements with len()

	fmt.Println("HAS ", len(mymap), " elements")

	// delete one element or clear all elements

	delete(mymap, "gay")
	clear(mymap)

	// use optional second return to check if a key was on the map

	_, present := mymap["gay"]

	fmt.Println("WAS IT PRESENT ?", present)

	// declare it with content

	map_con := map[string]int{"one": 1, "twe": 2}
	map_con2 := map[string]int{"one": 1, "twe": 2}
	if maps.Equal(map_con, map_con2) {
		fmt.Println("Are equial")
	}

}
