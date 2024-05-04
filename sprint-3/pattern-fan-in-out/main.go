package main

import (
	"log"
	"sync"
	"time"
)

type payload struct {
	name  string
	value int
}

func producer(name string, done <-chan struct{}, wg *sync.WaitGroup) <-chan payload {
	ch := make(chan payload)
	count := 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				log.Println(name, "completed")
				return
			case ch <- payload{name: name, value: count}:
				log.Println(name, "produced", count)
				count++
				time.Sleep(500 * time.Millisecond) //nolint:mnd // it's learning code
			}
		}
	}()

	return ch
}

func consumer(chs []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- payload) {
	for count, ch := range chs {
		go func() {
			defer wg.Done()
			log.Println("started consume of producer", count)
			for {
				select {
				case <-done:
					log.Println("consume of producer", count)
					return
				case v := <-ch:
					log.Println("consumer of producer", count, "got value", v.value, "from", v.name)
					fanIn <- v
				}
			}
		}()
	}
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2)                                 //nolint:mnd // it's learning code
	producers := make([]<-chan payload, 0, 3) //nolint:mnd // it's learning code
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))

	fanIn := make(chan payload)

	wg.Add(2) //nolint:mnd // it's learning code
	consumer(producers, done, &wg, fanIn)

	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-fanIn:
				log.Printf("fanIn got %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
