package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	log.Println(incrementGoroutine(1000)) //nolint:gomnd // it's example
}

func incrementGoroutine(count int) int {
	var wg sync.WaitGroup
	var num int64
	for range count {
		wg.Add(1)
		go increment(&num, 1, &wg)
	}

	wg.Wait()
	return int(num)
}

func increment(num *int64, add int64, wg *sync.WaitGroup) {
	defer wg.Done()

	atomic.AddInt64(num, add)
	time.Sleep(time.Microsecond)
}

func incrementWithoutDefer(num *int64, add int64) {
	atomic.AddInt64(num, add)
	time.Sleep(time.Microsecond)
}
