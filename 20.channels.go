package main

import (
	"fmt"
	"time"
)

// channels are pipes to connect to different goroutines

// send a vlaue form one goroutine and receive them in another

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

// pings can only send messages
func ping(pings chan<- string, msg string) {
	fmt.Println("ping func")
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("pong func")
	msg := <-pings
	pongs <- msg
}

func main() {
	// create a channel with:
	// 		make chan type
	messages := make(chan string) // sends strings

	// send a value with <- synthax
	go func() { messages <- "ping" }()

	// a send or a receive operation block until both sender and receiver are ready

	// so our main will get stuck until it receives something

	msg := <-messages
	fmt.Println("Received: ", msg)

	// BUFFERED CHANNELS

	// by default channels are unbuffered, this means that they will only accept a send if thjere is a corresponding receive

	// buffered accept a limited number of values without a corresponding receiver for those values

	// up to 3 values
	messages2 := make(chan string, 3)

	messages2 <- "buffered"
	messages2 <- "channel"
	messages2 <- "example"
	fmt.Println("\n\nBUFFERED:\n\n")
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	// CHANNEL SYNCHRONIZATION

	// WE CAN USE CHANNELS to synch work across goroutines

	// eg a blocking receive to wait for another goroutine to finish.

	// however if you have to wait for multiple goroutines, its better to use a WaitGroup (https://gobyexample.com/waitgroups)
	fmt.Println("\nCHANNEL SYNCH:\n")
	time.Sleep(time.Second)

	done := make(chan bool, 1)
	go worker(done)

	// block main thread here, until it receives it will not continue

	// without the next line the program coudl finsih even before the goroutine starts or before it ends
	<-done

	// CHANNEL DIRECTIONS

	// when a channel is a func param, you can specify if you only want it to send or receive
	// this increases type safety

	fmt.Println("\nCHANNEL DIRECTIONS:\n")
	time.Sleep(time.Second)

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
