package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -v ./error_types.go ./error_types_test.go
// === RUN   TestGetURL
// === RUN   TestGetURL/get_error_when_url_is_invalid
// === RUN   TestGetURL/get_error_when_url_is_invalid#01
// --- PASS: TestGetURL (0.00s)
// --- PASS: TestGetURL/get_error_when_url_is_invalid (0.00s)
// --- PASS: TestGetURL/get_error_when_url_is_invalid#01 (0.00s)
// PASS
// ok      command-line-arguments  0.391s

func TestGetURL(t *testing.T) {
	t.Run("get error when url is invalid", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer srv.Close()

		_, err := GetURL(srv.URL)
		if err == nil {
			t.Fatal("expected an error")
		}

		want := fmt.Sprintf("did not get 200 from %s, got %d", srv.URL, http.StatusTeapot)
		got := err.Error()
		if got != want {
			t.Fatalf("want %s, got %s", want, got)
		}
	})

	t.Run("get error when url is invalid", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer srv.Close()

		_, err := GetURLWithErrorType(srv.URL)
		if err == nil {
			t.Fatal("expected an error")
		}
		var got BadStatusError
		isStatusErr := errors.As(err, &got)
		if !isStatusErr {
			t.Fatalf("expected a BadStatusError, got %T", err)
		}
		want := BadStatusError{
			URL:    srv.URL,
			Status: http.StatusTeapot,
		}
		if !errors.Is(want, got) {
			t.Errorf("want %s, got %s", want, got)
		}
	})
}
