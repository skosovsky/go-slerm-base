package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v ./http_example.go ./http_example_test.go
// === RUN   TestRequest
// === RUN   TestRequest/success_request
// --- PASS: TestRequest (0.00s)
// --- PASS: TestRequest/success_request (0.00s)
// PASS
// ok      command-line-arguments  0.254s

func TestRequest(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(`{"Message": "ok"}`))
		if err != nil {
			panic(fmt.Errorf("cannot return http response: %w", err))
		}
	}))
	defer testServer.Close()

	t.Run("success request", func(t *testing.T) {
		req := Request{Value: "1"}
		result, err := sendRequest(req, testServer.URL)
		require.NoError(t, err)
		require.Equal(t, &Result{Message: "ok"}, result)
	})
}
