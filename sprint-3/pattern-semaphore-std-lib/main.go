package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func process(id int) {
	log.Printf("[%s]: running task %d\n", time.Now().Format("15:04:05"), id)
	time.Sleep(100 * time.Millisecond)
}

func main() {
	const n = 3
	const total = 100
	ctx := context.Background()

	sem := semaphore.NewWeighted(int64(n))

	for i := range total {
		if err := sem.Acquire(ctx, 1); err != nil {
			err = fmt.Errorf("failed to acquire semaphore: %w", err)
			log.Println(err)
			break
		}

		go func() {
			defer sem.Release(1)
			process(i)
		}()
	}

	if err := sem.Acquire(ctx, 1); err != nil {
		err = fmt.Errorf("failed to acquire semaphore: %w", err)
		log.Println(err)
	}
}
