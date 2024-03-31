package main

import (
	"log"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		log.Printf("Worker %d started job %d\n", id, job)

		// Do work and send result
		result := job * job
		results <- result

		log.Printf("Worker %d finished job %d\n", id, job)
	}
}

func doJobs(totalJobs int, totalWorkers int) {
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	// Start workers
	for w := 1; w <= totalWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Receive results
	for a := 1; a <= totalJobs; a++ {
		<-results // Place for result
	}
	close(results)
}

func main() {
	const totalJobs = 1000
	const totalWorkers = 5
	doJobs(totalJobs, totalWorkers)
}
