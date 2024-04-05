package main

import (
	"log"
	"sync"
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

	for j := 1; j <= totalJobs; j++ { // send jobs
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= totalJobs; a++ { // receive results
		<-results
	}
	close(results)
}
