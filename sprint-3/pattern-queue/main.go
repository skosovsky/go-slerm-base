package main

import (
	"log"
	"sync"
	"time"
)

func process(payload int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		log.Printf("Start processing of %d\n", payload)
		time.Sleep(500 * time.Millisecond) //nolint:gomnd // it's learning code
		log.Printf("Completed processing of %d\n", payload)

		<-queue
	}()
}

func main() {
	const n = 3
	const messages = 10
	var wg sync.WaitGroup

	log.Println("Queue of length M:", n)
	queue := make(chan struct{}, n)

	wg.Add(messages)

	for w := range messages {
		process(w, queue, &wg)
	}
}
