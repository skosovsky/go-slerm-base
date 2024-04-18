package main

import (
	"log"
	"sync"
	"time"
)

type empty struct{}

func calculate(num int) int {
	log.Printf("[%s] Calc for %d\n", time.Now().Format("15:04:05"), num)
	time.Sleep(1200 * time.Millisecond) //nolint:gomnd // it's learning code

	return num * 2 //nolint:gomnd // it's learning code
}

func main() {
	const dataSize = 10
	const semaphoreLimit = 3

	data := make([]int, 0, dataSize)
	for i := range dataSize {
		data = append(data, i+1)
	}

	results := make([]int, dataSize)
	semaphore := make(chan empty, semaphoreLimit)
	var wg sync.WaitGroup

	log.Printf("Before: %v\n", data)
	start := time.Now()

	for i, val := range data {
		wg.Add(1)
		go func() {
			defer wg.Done()
			semaphore <- empty{}
			results[i] = calculate(val)
			<-semaphore
		}()
	}
	wg.Wait()

	log.Printf("After: %v\n", results)
	log.Printf("Elapsed: %s\n", time.Since(start))
}
