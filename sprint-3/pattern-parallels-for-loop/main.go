package main

import (
	"log"
	"time"
)

func calculate(num int) int {
	time.Sleep(500 * time.Millisecond)
	return num * 2
}

func main() {
	const dataSize = 4

	data := make([]int, 0, dataSize)
	for i := range dataSize {
		data = append(data, i+10)
	}

	results := make([]int, dataSize)
	semaphore := make(chan int, dataSize)

	log.Printf("Before: %v\n", data)
	start := time.Now()

	for i, val := range data {
		go func() {
			results[i] = calculate(val)
			semaphore <- results[i]
		}()
	}
	for range dataSize {
		log.Printf("one calculation completed: %d\n", <-semaphore)
	}

	log.Printf("After: %v\n", results)
	log.Printf("Elapsed: %s\n", time.Since(start))
}
