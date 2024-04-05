package main

import (
	"fmt"
	"log"
	"time"
)

func echo(arg string, ch chan string) {
	ch <- arg // Send
}

func sendToChannelOnly(arg string, ch chan<- string) {
	ch <- arg // Send only
	// fmt.Println(<-ch) // Invalid operation: <-ch (receive from the send-only type chan<- string)
}

func receiveFromChannelOnly(ch <-chan string) {
	// ch <- "test" // Invalid operation: ch <- "test" (send to the receive-only type <-chan string)
	fmt.Println(<-ch) //nolint:forbidigo // it's learning code
}

func blockedChan() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	log.Println(<-ch)
}

func main() {
	// unbuffered channels
	var ch chan string
	fmt.Println(ch) //nolint:forbidigo // <nil>
	if ch != nil {  //nolint:govet // it's stub
		ch <- "Hello, DeadLock" // send to nil channel blocks forever // Panic: all goroutines are sleep - deadlock
		_ = <-ch                //nolint:gosimple // receive from nil channel blocks forever // Panic: all goroutines are sleep - deadlock
	}

	chInit := make(chan string)
	fmt.Println(chInit) //nolint:forbidigo // address memory

	go echo("Hello, Chan!", chInit)

	data := <-chInit  // Receive
	fmt.Println(data) //nolint:forbidigo // Hello, Chan!

	// buffered channels
	ch1 := make(chan string, 2) //nolint:gomnd // it's learning code
	go echo("value 1", ch1)
	go echo("value 2", ch1)

	val1 := <-ch1
	val2 := <-ch1
	fmt.Println(val1, val2) //nolint:forbidigo // value 1 value 2 || value 2 value 1

	close(ch1)
	// ch1 <- "Hello, Panic!" // send to close channel - panic // Panic: send on closed channel
	// Best practice - Who close channel? writer only!

	var c = make(chan int, 3)
	c <- 20
	c <- 10
	c <- 0
	close(c)

	// Receive from a closed channel returns the zero value.
	for range 5 {
		v, ok := <-c
		fmt.Printf("open?: %v, value %d\n", ok, v) //nolint:forbidigo // it's learning code
		// open?: true, value 20
		// open?: true, value 10
		// open?: true, value 0
		// open?: false, value 0
		// open?: false, value 0
	}

	var c2 = make(chan int, 3) //nolint:gomnd // it's learning code
	c2 <- 20
	c2 <- 10
	c2 <- 0
	close(c2)
	for v := range c2 {
		fmt.Printf("value %d\n", v) //nolint:forbidigo // it's learning code
		// value 20
		// value 10
		// value 0
	}

	var c3 = make(chan string)
	go receiveFromChannelOnly(c3)
	go sendToChannelOnly("test", c3)

	time.Sleep(time.Duration(1) * time.Second) // it's for goroutine

	blockedChan()
}
