package main

import (
	"fmt"
	"sync"
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

func main() {
	// fanIn
	i1 := generateWork([]int{0, 2, 6, 8})
	i2 := generateWork([]int{0, 2, 6, 8})

	out := fanIn(i1, i2)
	for value := range out {
		fmt.Println("fan in value:", value) //nolint:forbidigo // it's learning code
	}
}
