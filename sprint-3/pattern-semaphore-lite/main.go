package main

import (
	"log"
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
	sem := NewSemaphore(3) //nolint:gomnd // it's learning code

	for i := range 20 {

		go func() {
			sem.Acquire(1)
			defer sem.Release(1)
			time.Sleep(1 * time.Second)
			log.Println(i)
		}()
	}

	time.Sleep(15 * time.Second)
}
