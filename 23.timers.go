package main

import (
	"fmt"
	"time"
)

// when we want ot execute some code in the future or repeatedly (every X seconds)

// we will use built in timer and ticker (https://gobyexample.com/tickers)

func main() {

	// timers represent a single event in the future, after the wait, you will be notified in a channel
	timer1 := time.NewTimer(2 * time.Second)

	// block on the timer channel until it sends a value (end of countdown)
	<-timer1.C
	fmt.Println("Timer 1 ended")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// cancel the timer before it ends
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 STOPPED")
	}
	time.Sleep(2 * time.Second)

	// TICKERS

	// when you want ot do something in the future at regular intervals repeatedly
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	// it should tick 3 times before the done fires
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}
