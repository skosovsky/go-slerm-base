package main

import (
	"fmt"
	"time"
)

func selectForChan() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond) //nolint:gomnd // it's learning code
		one <- "One"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond) //nolint:gomnd // it's learning code
		two <- "Two"
	}()

	for range 2 {
		select {
		case result := <-one:
			fmt.Println("Received:", result) //nolint:forbidigo // it's learning code
		case result := <-two:
			fmt.Println("Received:", result) //nolint:forbidigo // it's learning code
		}
	}

	close(one)
	close(two)
}

func selectForChanDefault() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond) //nolint:gomnd // it's learning code
		one <- "One"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond) //nolint:gomnd // it's learning code
		two <- "Two"
	}()
	time.Sleep(1 * time.Second)
	select {
	case result := <-one:
		fmt.Println("Received:", result) //nolint:forbidigo // it's learning code
	case result := <-two:
		fmt.Println("Received:", result) //nolint:forbidigo // it's learning code
	default:
		fmt.Println("Received: nothing") //nolint:forbidigo // it's learning code
	}

	close(one)
	close(two)
}

func main() {
	selectForChan()
	selectForChanDefault()
	time.Sleep(1 * time.Second)
}
