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
	i := 1
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(ch)
				log.Println(name, "completed")
				return
			case ch <- payload{name: name, value: i}:
				log.Println(name, "produced", i)
				i++
				time.Sleep(500 * time.Millisecond) //nolint:gomnd // it's learning code
			}
		}
	}()

	return ch
}

func consumer(name string, chs []<-chan payload, done <-chan struct{}, wg *sync.WaitGroup, fanIn chan<- payload) {
	for i, ch := range chs {
		go func() {
			defer wg.Done()
			log.Println("started consume of producer", name, i)
			for {
				select {
				case <-done:
					log.Println("consumer", name, i, "completed")
					return
				case v := <-ch:
					log.Println("consumer", name, i, "got value", v.value, "from", v.name)
					fanIn <- v
				}
			}
		}()
	}
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(3)                                 //nolint:gomnd // it's learning code
	producers := make([]<-chan payload, 0, 3) //nolint:gomnd // it's learning code
	producers = append(producers, producer("Alice", done, &wg))
	producers = append(producers, producer("Jack", done, &wg))
	producers = append(producers, producer("Bob", done, &wg))

	fanIn1 := make(chan payload)
	fanIn2 := make(chan payload)

	wg.Add(3) //nolint:gomnd // it's learning code
	consumer("C1", producers, done, &wg, fanIn1)

	wg.Add(3) //nolint:gomnd // it's learning code
	consumer("C2", producers, done, &wg, fanIn2)

	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-fanIn1:
				log.Printf("fanIn1 got %v\n", v)
			case v := <-fanIn2:
				log.Printf("fanIn2 got %v\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	close(done)
	wg.Wait()
}
