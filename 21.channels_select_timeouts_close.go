package main

import (
	"fmt"
	"time"
)

// go lets you wait on multiple channel ops

func main() {

	// SELECT (wait for more than one channel, until all answer)
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "two"
	}()
	for range 2 {
		// wait for the 2 channels, executes the print when each one arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// CHANNEL TIMEOUTS

	// programs that connect to external resources or need to bound execution time use TIMEOUTS
	c3 := make(chan string, 1)
	c4 := make(chan string, 1)
	fmt.Println("\n\nTIMEOUTS:\n\n")
	time.Sleep(time.Second)

	// the channel isunbuffered so send operation is non blocking (BE AWARE)
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "RESULT 1"
	}()

	// Solution for a channel never read:
	select {
	case res := <-c1:
		fmt.Println("res", res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// in this case it will finish before the timeout
	go func() {
		time.Sleep(2 * time.Second)
		c4 <- "result 2"
	}()
	select {
	case res := <-c4:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}

	// NON BLOCKING CHANNEL OPERATIONS

	// basic sends and receives on channels are blocking

	// use select with default clause to implement non blocking sends, receives, and non-blocking multi-way selects
	fmt.Println("\n\nNON BLOCKING:\n\n")
	time.Sleep(time.Second)
	messages_ch := make(chan string)
	signals_ch := make(chan bool)

	select {
	case msg := <-messages_ch:
		fmt.Println("received messgae", msg)
	default:
		fmt.Println("no message received")
	}

	// provided that is has no receiver, it will not be sent
	msg := "hi man"
	select {
	case messages_ch <- msg:
		fmt.Println("sent msg, ", msg)
	default:
		fmt.Println("no message snet")
	}

	// multiple cases for a multi way non blocking slelect

	select {
	case msg := <-messages_ch:
		fmt.Println("received multi weay msg", msg)
	case sig := <-signals_ch:
		fmt.Println("receied signal", sig)
	default:
		fmt.Println("No activity")
	}

	// CLOSING CHANNELS

	// A CLOSING CHANNEL is when no more values will be sent --> communicate completion to the channel receivers
	fmt.Println("\n\nCLOSING CHANNELS:\n\n")
	time.Sleep(time.Second)

	// we send jobs from main to a worker goroutine with jobs channel
	// when there are no more jnos, we close jobs channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// in this form (2 values returned), more will be false when channel is closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				time.Sleep(time.Second)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job: ", j)
	}

	// close it with close()
	close(jobs)
	fmt.Println("main sent all jobs")

	<-done

	// reading FROM a closed channel is inmmediate
	// returns zero value of the underlying type
	// the second value is true is the received is ok
	// or false if a ZERO value is generated
	_, ok := <-jobs
	fmt.Println("RECEIVED ZERO VALUE?", !ok)
}
