package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	ctxBackground := context.Background()
	// context.TODO() - нужен для случаев, когда родительный контекст будет передан позднее (для удобной замены)

	ctx := context.WithValue(ctxBackground, "value", 1) //nolint:revive,staticcheck // it's learning code
	logger := logrus.New()
	ctx = context.WithValue(ctx, "logger", logger) //nolint:revive,staticcheck // it's learning code

	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()

	ctx, cancel2 := context.WithCancel(ctx) //nolint:ineffassign,staticcheck,wastedassign // it's learning code
	defer cancel2()

	http.HandleFunc("/", handleRequest)

	log.Println("Server is running...")
	err := http.ListenAndServe("localhost:4000", nil) //nolint:gosec // it's learning code
	if err != nil {
		panic(err)
	}
}

// curl localhost:4000
func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Handle started")
	ctx := r.Context()
	requestID := ctx.Value("request-id")
	if requestID == nil {
		ctx = context.WithValue(ctx, "request-id", uuid.New().String()) //nolint:revive,staticcheck // it's learning code
	}

	select {
	case <-time.After(5 * time.Second): //nolint:gomnd // it's learning code
		_, _ = fmt.Fprint(w, "Response from the server")

	// Handling request cancellation
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("Context timed out:", err) // Error: context canceled
	}

	log.Println("Handle complete")
}
