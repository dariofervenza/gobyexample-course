package main

import (
	"fmt"
	"os"
	s "strings"
)

// std lib: strings

// contains useful string-related funcs

var p = fmt.Println

func main() {

	// take a look on https://pkg.go.dev/strings
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"test", "sentece", "to", "join"}, "-"))
	p("Repeat: ", s.Repeat("test", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Split: ", s.Split("a-b-c-d", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))

	// ------- STRING FORMATTING ------

	// go follows the printf tradition for formatting

	type point struct {
		x, y int
	}

	pfort := point{1, 2}

	fmt.Printf("struct1: %v\n", pfort)

	// there the %v is a verb used to format
	// %v prints the instance of the struct

	// %+v --> include struct field names

	fmt.Printf("struct2: %+v\n", pfort)

	// %#v prints a go synthax representation of the value --> eg. the code that would produce that value

	fmt.Printf("struct3: %#v\n", pfort)

	// Print the type of a variable

	fmt.Printf("struct4 type: %T\n", pfort)

	// bool formatting: %t
	fmt.Printf("boolean: %t\n", true)

	// integers
	fmt.Printf("int: %d\n", 55)

	// binary representation of int

	fmt.Printf("int in binary: %b\n", 55)

	// the character that corresponds to an integer

	fmt.Printf("%d is the char: %c\n", 33, 33)

	// hexagonal encoding

	fmt.Printf("hexagonal: %x\n", 55)

	// float formatting

	fmt.Printf("float: %f\n", 76.5)

	// %e and %E are scientific notation
	fmt.Printf("float notation: %e\n", 76.5)
	fmt.Printf("float notation (ver 2): %E\n", 76.5)

	// strings
	fmt.Printf("string: %s\n", "halo")

	// string with double quotes
	fmt.Printf("string double quotes: %q\n", "halo")

	// strings can be also printed with %x for the hexadecimal representation --> two output characters per byte of input
	fmt.Printf("str3 hexadecimal: %x\n", "hex this")

	// print the representation of a pointer
	fmt.Printf("pointer: %p\n", &p)

	// print integers with a fixed width --> 1 ---> "     1"  (five spaces plus 1 char)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// FLOATS OF FIXED WIDTH AND FIXED NUMBER OF DECIMALS
	fmt.Printf("width2 floats: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// justify them with f (flag)
	fmt.Printf("width3 justified: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // -->> the spaces after the number

	// control string widths (align them in tables)

	fmt.Printf("width4 aligned strings: |%6s|%6s|\n", "foo", "b")

	// they also have the f option
	fmt.Printf("width5 string with flag (justified): |%-6s|%-6s|\n", "foo", "b")

	// ------------- EXPLANATION ------------
	// What the Printf actually does?

	// it does a Sprintf operation which returns the formatted strings

	sFormatted := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(sFormatted)

	// PrinfF writes directly to stdout (os.Stdout) but you get the idea

	// --- dump formatted string into a writer -----------------

	// Format + PRINT TO A WRITER --> io.Writer instead of os.Stdout

	fmt.Fprintf(os.Stderr, "io: an %s\n", "errror snet to stderr")

}
