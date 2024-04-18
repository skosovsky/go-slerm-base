package main

import (
	"log"
	"sync"
)

type mutexType chan struct{}
type mutex struct {
	s mutexType
}

func newMutex() mutex {
	return mutex{s: make(mutexType, 1)}
}

func (m mutex) Lock() {
	e := struct{}{}
	m.s <- e
}

func (m mutex) Unlock() {
	<-m.s
}

func main() {
	const count = 1000

	customMutex := newMutex()
	counter := 0
	var wg sync.WaitGroup

	wg.Add(count)
	for range count {
		go func() {
			defer wg.Done()
			customMutex.Lock()
			defer customMutex.Unlock()
			counter++
		}()
	}

	wg.Wait()
	log.Printf("Mutex counter: %d\n", counter)
}
