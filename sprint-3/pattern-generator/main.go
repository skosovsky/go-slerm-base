package main

import (
	"fmt"
)

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	ch := generator()

	for range 5 {
		value := <-ch
		fmt.Println("generator value:", value) //nolint:forbidigo // it's learning code
	}
}
