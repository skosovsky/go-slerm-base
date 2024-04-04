package main

import (
	"crypto/md5"
	"encoding/hex"
	"sync"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	go func() {
		defer w.wg.Done()
		for task := range w.tasks {
			hash := md5.Sum([]byte(task))
			w.out <- hex.EncodeToString(hash[:])
		}
	}()
}

func main() {
	const totalJobs = 100
	const totalWorkers = 5
	tasks := make(chan string, totalJobs)
	out := make(chan string, totalJobs)

	var wg sync.WaitGroup
	for w := 1; w <= totalWorkers; w++ { // prepare workers
		worker := NewWorker(tasks, &wg, out)
		wg.Add(1)
		worker.Run()
	}

	for j := 1; j <= totalJobs; j++ { // send jobs
		tasks <- string(rune(j))
	}
	close(tasks)

	wg.Wait()
	for a := 1; a <= totalJobs; a++ { // receive results
		println(<-out)
	}
	close(out)
}
