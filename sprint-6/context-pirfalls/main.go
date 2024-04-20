package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type oneParam string
type otherParam string

func main() {
	// context values overriding
	ctx := context.Background()               // root of all derived contexts
	ctx = context.WithValue(ctx, "param", 10) //nolint:revive,gomnd,staticcheck // it's learning code
	log.Println(ctx.Value("param"))           // 10
	ctx = context.WithValue(ctx, "param", 20) //nolint:revive // it's learning code
	log.Println(ctx.Value("param"))           // 20

	ctx2 := context.Background()
	var param1 oneParam = "param"
	ctx2 = context.WithValue(ctx2, param1, 10)
	var param2 otherParam = "param"
	ctx2 = context.WithValue(ctx2, param2, 20)
	log.Println(ctx2.Value(param1)) // 10
	log.Println(ctx2.Value(param2)) // 20

	// advisory cancellation
	ctx3, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx3.Done():
			log.Println("Get cancel signal")
			time.Sleep(3 * time.Second)
			log.Println("cancelled")
		}
	}()
	cancel()

	// explicit better than implicit
	ctx4 := context.Background()
	ctx4 = context.WithValue(ctx4, "importantParam", 123) //nolint:revive,gomnd // it's learning code
	myMethodImplicit(ctx4)

	// better
	myMethodExplicit(ctx4, 123)

	// explicit context value usage
	// addRequestID
	// useRequestID

	// context cancel propagation
	ctx5, cancel5 := context.WithCancel(context.Background())
	go cancelLabelMethod(ctx5)
	cancel5()

	networkRequest()

	time.Sleep(5 * time.Second)
}

func addRequestID(w http.ResponseWriter, r *http.Request, next http.Handler) {
	ctx := context.WithValue(r.Context(), "request-id", r.Header.Get("request-id")) //nolint:revive,staticcheck // it's learning code
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
}

func useRequestID(_ http.ResponseWriter, r *http.Request, _ http.Handler) {
	requestID := r.Context().Value("request-id")
	log.Println("send request to other service", requestID)
}

func myMethodImplicit(ctx context.Context) {
	val := ctx.Value("importantParam")
	log.Println("importantParam", val) // importantParam 123
}

func myMethodExplicit(_ context.Context, importantParam int) {
	log.Println("importantParam", importantParam) // importantParam 123
}

func cancelLabelMethod(ctx context.Context) {
	ctx = context.WithValue(ctx, "param", "value") //nolint:revive,staticcheck // it's learning code
	cancelLabelMethod2(ctx)
}

func cancelLabelMethod2(ctx context.Context) {
	if <-ctx.Done(); true {
		log.Println("cancelled inner context two")
	}
}

func networkRequest() {
	// Get "https://example.com": dial tcp: lookup example.com: i/o timeout
	const timeout = 1 * time.Second

	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{ //nolint:exhaustruct // it' practice
				Timeout: timeout,
			}).DialContext,
		},
	}

	// Get "https://example.com": context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
	if err != nil {
		log.Println(err)
		return
	}

	rsp, err := client.Do(req)
	if rsp == nil {
		log.Printf("nil response\n")
		return
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(rsp.Body)

	var e net.Error
	if errors.As(err, &e) && e.Timeout() {
		log.Printf("do request timeout: %s\n", err)
		return
	} else if err != nil {
		log.Printf("cannot do request: %s\n", err)
		return
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("body: %s\n", body)
}
