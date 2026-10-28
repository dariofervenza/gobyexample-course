package main

import (
	"fmt"
	"time"
)

// when you need the number of seconds, ms or ns from the unix epoch (https://en.wikipedia.org/wiki/Unix_time)

func main() {
	now := time.Now()
	fmt.Println("Now is: ", now)

	// unix seconds
	fmt.Println("Unix secs: ", now.Unix())
	// unix miliseconds
	fmt.Println("Unix milisecs: ", now.UnixMilli())
	// unix nano seconds
	fmt.Println("Unix nanosecs: ", now.UnixNano())

	// create a Time struct from unix seconds
	fmt.Println("Time from Unix secs: ", time.Unix(now.Unix(), 0))
	// CREATE a Time struct from unix nano seconds
	fmt.Println("Time from Unix secs: ", time.Unix(0, now.UnixNano()))

}
