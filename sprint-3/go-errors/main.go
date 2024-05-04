package main

import (
	"fmt"
	"sync"
	"time"
)

func doIt(workerID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[%v] is running\n", workerID) //nolint:forbidigo // it's learning code
	time.Sleep(3 * time.Second)               //nolint:mnd // it's learning code
	fmt.Printf("[%v] is done\n", workerID)    //nolint:forbidigo // it's learning code
}

func Search(query string) string {
	time.Sleep(time.Second)
	return "Result of " + query
}

func First(query string, replicas ...func(string) string) string {
	ch := make(chan string) // another goroutines make leak memory, because don't finished work
	// c := make(chan string, len(replaces)) // firs solution
	searchReplica := func(i int) { ch <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}

	return <-ch
}

func FirstWithSelect(query string, replicas ...func(string) string) string { // second solution
	ch := make(chan string, 1)
	searchReplica := func(i int) {
		select {
		case ch <- replicas[i](query):
		default:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-ch
}

func FirstWithCancel(query string, replicas ...func(string) string) string {
	ch := make(chan string)
	done := make(chan struct{})
	defer close(done)
	searchReplica := func(i int) {
		select {
		case ch <- replicas[i](query):
		case <-done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}

	return <-ch
}

func main() {
	// Use WaitGroup for wait finished goroutines
	var wg sync.WaitGroup
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, &wg)
	}

	wg.Wait()                // It's good
	fmt.Println("all done!") //nolint:forbidigo // it's learning code

	// First result, maybe leak resources
	s := First("test", Search, Search)
	fmt.Println(s) //nolint:forbidigo // it's learning code

	s2 := FirstWithSelect("test", Search, Search)
	fmt.Println(s2) //nolint:forbidigo // it's learning code

	s3 := FirstWithCancel("test", Search, Search)
	fmt.Println(s3) //nolint:forbidigo // it's learning code

	// Popular error fix in Go 1.22
	var wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			fmt.Printf("%d", i) //nolint:forbidigo // it's learning code
			wg2.Done()
		}()
	}
	wg2.Wait()
}
