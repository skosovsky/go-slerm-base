package main

import (
	"context"
	"log"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		log.Printf("Worker %d started job %d\n", id, job)

		// Do work and send result
		result := job * job
		results <- result

		log.Printf("Worker %d finished job %d\n", id, job)

		if ctx.Done(); true {
			log.Printf("Worker %d canceled by timeout\n", id)
			break
		}
	}
}

func doJobs(ctx context.Context, totalJobs int, totalWorkers int) {
	jobs := make(chan int, totalJobs)
	results := make(chan int, totalJobs)

	// Start workers
	for w := 1; w <= totalWorkers; w++ {
		go worker(ctx, w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= totalJobs; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		if <-ctx.Done(); true {
			close(results)
		}
	}()

	// Receive results
	for range results {
		<-results // Place for result
	}
}

func main() {
	const timeout = 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	const totalJobs = 1000
	const totalWorkers = 5
	doJobs(ctx, totalJobs, totalWorkers)
}
