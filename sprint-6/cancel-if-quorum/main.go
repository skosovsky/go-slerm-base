package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func getRequest(ctx context.Context, url string) error {
	go func() {
		if <-ctx.Done(); true {
			log.Println("cancelled")
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	rsp, err := http.DefaultClient.Do(req) //nolint:bodyclose // it's closed
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(rsp.Body)

	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", rsp.Status)
	}

	return nil
}

func doManyRequest(ctx context.Context, count int64) bool {
	var isQuorum = make(chan struct{})
	var quorumCount int64
	var mu sync.Mutex

	go func() {
		if <-ctx.Done(); true {
			log.Println("ctx cancelled")
		}
	}()

	for range count {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			log.Println("lo")
			err := getRequest(ctx, "https://example.com")
			if err == nil {
				quorumCount++
				log.Println(quorumCount)
			}

			if quorumCount > count/2 {
				isQuorum <- struct{}{}
				ctx.Done()
			}
		}()
	}
	select {
	case <-isQuorum:
		return true
	case <-ctx.Done():
		return false
	}
}

func main() {
	const timeout = 2 * time.Second
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	log.Println(doManyRequest(ctx, 5)) //nolint:mnd //it's example
}
