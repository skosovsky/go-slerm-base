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
	//  generator
	ch := generator()

	for range 5 { //nolint:typecheck // it's ok for 1.22
		value := <-ch
		fmt.Println("generator value:", value) //nolint:forbidigo // it's learning code
	}

}
