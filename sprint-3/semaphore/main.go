package main

import (
	"fmt"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func (s Semaphore) Acquire(n int) {
	for range n {
		s <- struct{}{}
	}
}

func (s Semaphore) Release(n int) {
	for range n {
		<-s
	}
}

func main() {
	sem := NewSemaphore(10) //nolint:gomnd // it's learning code

	for i := 0; i < 20; i++ {
		go func() {
			sem.Acquire(1)
			defer sem.Release(1)
			fmt.Println(i) //nolint:forbidigo // it's learning code
		}()
	}

	time.Sleep(time.Second)
}
