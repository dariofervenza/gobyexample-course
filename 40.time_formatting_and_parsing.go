package main

import (
	"fmt"
	"time"
)

// in order to print time into a formatted string

// or to get a Time struct from a string  we use the following code

func main() {
	p := fmt.Println

	now := time.Now()

	// format time -> string

	p("Time to RFC3339: ", now.Format(time.RFC3339))

	// time string to time struct

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

	p("WE got from a str: ", t1)

	// HOWEVER: Parse and format work by sending an example
	// (no need for %y-%M-%d and similar, with a simple example it inferes the format)
	p(now.Format("3:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// alternatively you could use string formatting and EXTRACTED COMPONENTS
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// if the sent format is incoorect, it will return an err
	_, err := time.Parse("Mon Jan _2 15:04:05 2006", "8:41PM")
	p(err)

}
