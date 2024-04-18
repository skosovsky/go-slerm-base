package main

import (
	"fmt"
	"sync"
	"time"
)

func process() {
	time.Sleep(time.Millisecond)
	fmt.Println("process") //nolint:forbidigo // it's learning code
}

func processWithWG(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond)
	fmt.Println("process with wait group") //nolint:forbidigo // it's learning code
}

func waitGroupWithGoFunc() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		process()
	}()

	wg.Wait()
}

func waitGroupWithExternalFunc() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		processWithWG(&wg)
	}()

	wg.Wait()
}

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	defer c.m.Unlock()
	c.value += n
}

func mutexLock() {
	var wg sync.WaitGroup
	counter := Counter{
		m:     sync.Mutex{},
		value: 0,
	}

	for range 100 {
		wg.Add(1)
		go counter.Update(10, &wg) //nolint:gomnd // it's learning code
	}
}

type RWCounter struct {
	m     sync.RWMutex
	value int
}

func (r *RWCounter) Update(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	r.m.Lock()
	defer r.m.Unlock()
	r.value += n
}

func (r *RWCounter) Value() int {
	r.m.RLock()
	defer r.m.RUnlock()
	v := r.value

	return v
}

func mutexRWLock() {
	var wg sync.WaitGroup

	counter := RWCounter{
		m:     sync.RWMutex{},
		value: 0,
	}

	for range 100 {
		wg.Add(1)
		go counter.Update(10, &wg) //nolint:gomnd // it's learning code
	}

	wg.Wait()
	fmt.Printf("Result is %d\n", counter.Value()) //nolint:forbidigo // it's learning code
}

func syncMap() {
	var wg sync.WaitGroup
	var dataMap sync.Map // dataMap not nil, but don't copy

	wg.Add(10) //nolint:gomnd // it's learning code

	for i := 1; i <= 5; i++ {
		go func(key int) {
			value := fmt.Sprintf("value %d", key)
			fmt.Println("Writing:", value) //nolint:forbidigo // it's learning code

			dataMap.Store(key, value)
			wg.Done()
		}(i)
	}

	for i := 1; i <= 5; i++ {
		go func(key int) {
			value, ok := dataMap.Load(key)
			if ok {
				fmt.Println("Reading:", value) //nolint:forbidigo // it's learning code
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func doDataRace() {
	channel := make(chan bool)
	data := make(map[string]string)

	go func() {
		data["1"] = "a" // First conflicting access
		channel <- true
	}()

	data["2"] = "b" // Second conflicting access
	<-channel       // Analog Wait Group, locking main by read

	for key, value := range data {
		fmt.Println(key, value) //nolint:forbidigo // it's learning code
	}
}

func main() {
	waitGroupWithGoFunc()
	waitGroupWithExternalFunc()
	mutexLock()
	mutexRWLock()
	syncMap()
	doDataRace() // go run -race . // WARNING: DATA RACE
}
