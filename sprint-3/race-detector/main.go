package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range 10000 {
		mu.Lock()
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mu.Unlock()
			count++
		}()
	}

	wg.Wait()
	fmt.Println(count) //nolint:forbidigo // it's learning code
}
