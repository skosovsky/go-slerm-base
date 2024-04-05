package main

import (
	"log"
	"math"
)

func generateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()

	return ch
}

func pipeline(f func(i int) (int, bool), in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			if value, ok := f(v); ok {
				out <- value
			}
		}
	}()

	return out
}

func filterOdd(num int) (int, bool) {
	if num%2 == 0 {
		return num, true
	}
	return 0, false
}

func square(num int) (int, bool) {
	value := math.Pow(float64(num), 2)
	return int(value), true
}

func half(num int) (int, bool) {
	return num / 2, true
}

func main() {
	in := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out := pipeline(filterOdd, in)
	out = pipeline(square, out)
	out = pipeline(half, out)

	for value := range out {
		log.Println(value)
	}
}
