package main

import (
	"fmt"
	"unicode/utf8"
)

// in go the concept of a character is called a rune
// a string is a read only slice of bytes
// a rune is an integer that represent the character in UTF-8

func main() {

	const s = "สวัสดี"

	// strings are equivalent to []byte so we can use len() in them
	fmt.Println("Len of s: ", len(s))

	for i := 0; i < len(s); i++ {
		// print the hexadecimal value of the byte
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// count how many runes are in a string
	// RuneCountInString decodes eahc UTF-8 char
	// some UTF-8N code points can spawn multiple bytes so the result of the count is surprising
	s1 := "im a s s s"
	fmt.Println("Len s: ", len(s), "\nlen s1: ", s1)
	fmt.Println("Rune count s: ", utf8.RuneCountInString(s))
	fmt.Println("Rune count s1: ", utf8.RuneCountInString(s1))

	// a range loop decodes each UTF-8 rune along its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// you can do the same with the utf8.DecodeRuneInString function
	fmt.Println("\nUsing DecodeRuneInString\n")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examine_rune(runeValue)
	}
}

func examine_rune(r rune) {
	// NOTICE the simple commas ''
	// with "t" we are comparing different types -> rune != string
	// values enclosed in '' are rune literals
	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'ส' {
		fmt.Println("\nfound so sua\n")
	}
}
