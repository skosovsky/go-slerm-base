package main

import (
	"fmt"
	"math"
)

func generateWork(works []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, work := range works {
			ch <- work
		}
	}()

	return ch
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2) //nolint:gomnd // it's learning code
			out <- int(value)
		}
	}()

	return out
}

func main() {
	// pipeline
	pipeIn := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out1 := filter(pipeIn) // Filter add numbers
	out1 = square(out1)    // Square the input

	for value := range out1 {
		fmt.Println("pipeLine value:", value) //nolint:forbidigo // it's learning code
	}
}
