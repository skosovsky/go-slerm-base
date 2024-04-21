package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	const timeout = 100 * time.Millisecond
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	go func() {
		if <-ctx.Done(); true {
			log.Println("cancelled")
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com/", nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	rsp, err := http.DefaultClient.Do(req) //nolint:bodyclose // it's closed
	if err != nil {
		log.Println(err)
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(rsp.Body)

	log.Println(rsp.StatusCode)
}
