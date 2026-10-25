package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// go supports regex with the stdlib: regexp

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	// it will return true
	fmt.Println(match)

	// create a pattern with Compile
	// r will be an optimized regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// the regexp struct has many methods

	fmt.Println(r.MatchString("peach"))

	// find in a string (only first match):
	fmt.Println(r.FindString("peach punch")) // prints "peach"

	// get the indexes of the first match

	fmt.Println("idx:", r.FindStringIndex("peach punch")) // [0, 5]

	// ---- SUB MATCH ----
	// we use submatch  to get first match and the groups ([a-z])

	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// get indexes of first match and group (submatch)
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// ---- ALL MATCHES ---
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// THE All variants are available in the other functions
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// limit the max number of coincidences (eg 2)
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// ---- bytes instead of str --------------------------
	// send strings as byte slice (drop string from func signature)
	fmt.Println(r.Match([]byte("peach"))) // true

	// ------------- ensure REGEX PATTERN WORKS -------------
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) // p([a-z]+)ch

	// ---- REGEX REPLACE ----
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// ----- FUNCTIONAL REPLACE ----
	// replace matched text with a function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	// or with strings
	in1 := "a peach"
	out1 := r.ReplaceAllStringFunc(in1, strings.ToUpper)
	fmt.Println(out1)

	// hence if the function is DoSomethingString --> uses strings
	// if the function is DoSomethingStringFunc   ---> uses a function
	// if the function is DoSomething ----------------> uses byte slice
}
