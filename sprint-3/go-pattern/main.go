package main

import (
	"fmt"
	"math"
	"sync"
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

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	var out = make(chan int)

	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()

			for value := range ch {
				out <- value
			}

		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
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
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func main() {
	//  generator
	ch := generator()

	for range 5 {
		value := <-ch
		fmt.Println("generator value:", value) //nolint:forbidigo // it's learning code
	}

	// fanIn
	i1 := generateWork([]int{0, 2, 6, 8})
	i2 := generateWork([]int{0, 2, 6, 8})

	out := fanIn(i1, i2)
	for value := range out {
		fmt.Println("fan in value:", value) //nolint:forbidigo // it's learning code
	}

	// pipeline
	pipeIn := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out1 := filter(pipeIn) // Filter add numbers
	out1 = square(out1)    // Square the input

	for value := range out1 {
		fmt.Println("pipeLine value:", value) //nolint:forbidigo // it's learning code
	}

	// worker pool
	// TODO: продолжить
}
