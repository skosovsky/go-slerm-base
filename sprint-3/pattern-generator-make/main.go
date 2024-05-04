package main

import (
	"log"
	"sync"
	"time"
)

func makeGenerator(done <-chan struct{}, wg *sync.WaitGroup) <-chan int {
	var i = 0
	ch := make(chan int, 1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				log.Println("done")
				return
			default:
				time.Sleep(250 * time.Millisecond) //nolint:mnd // it's learning code
				ch <- i
				i++
			}
		}
	}()

	return ch
}

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wg.Add(2) //nolint:mnd // it's learning code

	ch := makeGenerator(done, &wg)

	go func() {
		defer wg.Done()
		for v := range ch {
			log.Println("value:", v)
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
