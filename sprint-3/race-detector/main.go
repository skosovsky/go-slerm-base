package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var m sync.Mutex
	var wg sync.WaitGroup

	for range 10000 {
		m.Lock()
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer m.Unlock()
			count++
		}()
	}

	wg.Wait()
	fmt.Println(count) //nolint:forbidigo // it's learning code
}
