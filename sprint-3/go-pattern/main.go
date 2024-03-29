package main

import (
	"fmt"
	"math"
	"sync"
)

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func generateWork(works []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, work := range works {
			ch <- work
		}
	}()

	return ch
}

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	var out = make(chan int)

	wg.Add(len(inputs))

	for _, input := range inputs {
		go func(ch <-chan int) {
			defer wg.Done()
			for value := range ch {
				out <- value
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2) //nolint:gomnd // it's learning code
			out <- int(value)
		}
	}()

	return out
}

func worker(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup

	for job := range jobs {
		wg.Add(1)

		go func(job int) {
			defer wg.Done()

			fmt.Printf("Worker %d started job %d\n\n", id, job) //nolint:forbidigo // it's learning code

			// Do work and send results
			result := job * job
			results <- result

			fmt.Printf("Worker %d finished job %d", id, job) //nolint:forbidigo // it's learning code
		}(job)
	}

	wg.Wait()
}

func main() {
	//  generator
	ch := generator()

	for range 5 { //nolint:typecheck // it's ok for 1.22
		value := <-ch
		fmt.Println("generator value:", value) //nolint:forbidigo // it's learning code
	}

	// fanIn
	i1 := generateWork([]int{0, 2, 6, 8})
	i2 := generateWork([]int{0, 2, 6, 8})

	out := fanIn(i1, i2)
	for value := range out {
		fmt.Println("fan in value:", value) //nolint:forbidigo // it's learning code
	}

	// pipeline
	pipeIn := generateWork([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out1 := filter(pipeIn) // Filter add numbers
	out1 = square(out1)    // Square the input

	for value := range out1 {
		fmt.Println("pipeLine value:", value) //nolint:forbidigo // it's learning code
	}

	// worker pool
	const totalJobs = 10
	const totalWorkers = 5
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	for w := 1; w <= totalWorkers; w++ { // prepare workers
		go worker(w, jobs, results)
	}

	for j := 1; j <= totalJobs; j++ { // send jobs
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= totalJobs; a++ { // receive results
		<-results
	}
	close(results)
}
