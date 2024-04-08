package main

import (
	"log"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup

	for job := range jobs {
		wg.Add(1)

		go func(job int) {
			defer wg.Done()

			log.Printf("Worker %d started job %d\n\n", id, job) //nolint:forbidigo // it's learning code

			// Do work and send results
			result := job * job
			results <- result

			log.Printf("Worker %d finished job %d", id, job) //nolint:forbidigo // it's learning code
		}(job)
	}

	wg.Wait()
}

func main() {
	// worker pool
	const totalJobs = 1000
	const totalWorkers = 5
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	for w := 1; w <= totalWorkers; w++ { // prepare workers
		go worker(w, jobs, results)
	}

	for j := range totalJobs { // send jobs
		jobs <- j
	}
	close(jobs)

	for range totalJobs { // receive results
		<-results
	}
	close(results)
	time.Sleep(time.Second)
}
