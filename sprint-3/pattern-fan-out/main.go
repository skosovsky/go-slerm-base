package main

import (
	"log"
	"sync"
)

func generateWork(work []int, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)

	go func() {
		defer wg.Done()
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
		log.Println("All data written")
	}()

	return ch
}

func fanOut(in <-chan int, name string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()

		for data := range in {
			log.Println(name, "processed", data)
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	work := []int{1, 2, 3, 4, 5, 6, 7, 8}

	wg.Add(1)
	in := generateWork(work, &wg)

	wg.Add(3)
	fanOut(in, "Alice", &wg)
	fanOut(in, "Jack", &wg)
	fanOut(in, "Bob", &wg)

	wg.Wait()
}
