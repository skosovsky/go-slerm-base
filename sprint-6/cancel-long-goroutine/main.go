package main

import (
	"context"
	"log"
	"time"
)

func main() {
	const timeout = 5 * time.Second
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go func(ctx context.Context) {
		var count int
		for {
			select {
			case <-ctx.Done():
				log.Println(count)
				return
			default:
				time.Sleep(3 * time.Second) //nolint:gomnd // it's learning code
				count++
			}
		}
	}(ctx)

	time.Sleep(7 * time.Second) //nolint:gomnd // it's learning code
}
