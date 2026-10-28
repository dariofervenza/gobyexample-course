package main

import (
	"fmt"
	"time"
)

// go support for time and durations

func main() {
	p := fmt.Println

	// get current time
	now := time.Now()
	p("Now is: ", now) // 2026-10-28 08:52:59.749964851 +0100 CET m=+0.000021344

	// create a time strcut by giving the year, month day, hours,. ..
	// time is always linked to a location, hence as timezone
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p("Created as date: ", then) // the previous was UTC +0100, this is UTC +0000

	// extract the components of a time
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// also the weekday
	p(then.Weekday())

	// check if one date occured before, after or at the same moment:
	p("Check is before: ", then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// -------- DURATION struct: time differences ----------------
	diff := now.Sub(then)
	p("Difference is a duration struct: ", diff) // 148523h24m50.898340074s

	// transform the difference to the units you need
	p("Difference in hours (total): ", diff.Hours())
	p("Difference in minutes: ", diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add method, to advance a date
	p("Then + diff is now: ", then.Add(diff))
	p("Then - diff is even older: ", then.Add(-diff)) // 1992-12-08 09:06:58.355699562 +0000 UTC

	//  SUMMARY: .Sub() gets a Duration whereas .Add() gets a new Time struct

}
