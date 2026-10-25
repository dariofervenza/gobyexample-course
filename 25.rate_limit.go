package main

import (
	"fmt"
	"time"
)

// rate limiting is a mechanism for controlling resource utilization and maintain quality of serv (https://en.wikipedia.org/wiki/Rate_limiting).
// we can rate limit wirg goroutines, channels and tickers

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// thev limiter channel will receive a value every 200 ms, we limit ourselves to 1 request every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		// USE THE BLOCKING CHANNEL TO LIMIT
		<-limiter
		fmt.Println("Limiter freed request", req, time.Now())
	}

	// BURSTS REQUESTS: send 3 request at the same time, then one each x seconds (ACCUMULATE MAX 3 TO RUN)

	time.Sleep(time.Second)
	fmt.Println()
	// allow short bursts of requests in our rate limiting while preserving the overall ratel imit
	// allow buursts of up to 3 events

	burstyLimiter := make(chan time.Time, 3)

	// fill the limiter
	for range 3 {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(950 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// EXECUTE WORK BASED ON BURSTY LIMITER
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request form bursty", req, time.Now())
	}

}
