package main

import (
	"log"
	"sync"
	"time"
)

func small() {
	type smallStruct struct {
		value int
	}

	ch := make(chan smallStruct, 10000000*2) //nolint:gomnd // it's learning code
	var wg sync.WaitGroup
	var duration time.Duration
	var startTime time.Time

	startTime = time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 10000000 {
			ch <- smallStruct{value: 1000} //nolint:gomnd // it's learning code
			ch <- smallStruct{value: -1000}
		}
		close(ch)
	}()
	wg.Wait()
	duration = time.Since(startTime)
	log.Println("[small] Transfer to chan:", duration)

	startTime = time.Now()
	var sum int
	for val := range ch {
		sum += val.value
	}
	duration = time.Since(startTime)
	log.Println("[small] Count and transfer from chan:", duration, sum)
}

func big() {
	type bigStruct struct {
		name  string
		email string
		age   int
		value int
	}

	ch := make(chan bigStruct, 10000000*2) //nolint:gomnd // it's learning code
	var wg sync.WaitGroup
	var duration time.Duration
	var startTime time.Time

	startTime = time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 10000000 {
			ch <- bigStruct{name: "Sfdfsdfsdfsdf", email: "sg@mail.ru", age: 54, value: 1000}  //nolint:gomnd // it's learning code
			ch <- bigStruct{name: "Sfsdfsdfsfsdf", email: "sg@mail.ru", age: 54, value: -1000} //nolint:gomnd // it's learning code
		}
		close(ch)
	}()
	wg.Wait()
	duration = time.Since(startTime)
	log.Println("[BIG] Transfer to chan:", duration)

	startTime = time.Now()
	var sum int
	for val := range ch {
		sum += val.value
	}
	duration = time.Since(startTime)
	log.Println("[BIG] Count and transfer from chan:", duration, sum)
}

func main() {
	time.Sleep(time.Second)
	small()
	time.Sleep(time.Second)
	big()
}
