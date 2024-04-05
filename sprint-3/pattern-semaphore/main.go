package main

import (
	"log"
	"time"
)

type semaphore chan struct{}

func newSemaphore(n int) semaphore {
	return make(semaphore, n)
}

func (s semaphore) acquire(n int) {
	for range n {
		s <- struct{}{}
	}
}

func (s semaphore) release(n int) {
	for range n {
		<-s
	}
}

func process(id int) {
	log.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(time.Second)
}

func main() {
	const n = 3
	const total = 10

	sem := newSemaphore(n)
	done := make(chan bool)

	for i := range total {
		sem.acquire(1)
		go func() {
			defer sem.release(1)
			process(i)
			if i == total-1 {
				done <- true
			}
		}()
	}

	<-done
}
