package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetURL(url string) (string, error) {
	res, err := http.Get(url) //nolint:gosec,bodyclose,noctx // it's learning code
	if err != nil {
		return "", fmt.Errorf("error while getting url: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("did not get 200 from %s, got %d", url, res.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading body: %w", err)
	}

	return string(body), nil
}

type BadStatusError struct {
	URL    string
	Status int
}

func (b BadStatusError) Error() string {
	return fmt.Sprintf("Bad status for URL %q: %d ", b.URL, b.Status)
}

func GetURLWithErrorType(url string) (string, error) {
	res, err := http.Get(url) //nolint:gosec,bodyclose,noctx // it's learning code
	if err != nil {
		return "", fmt.Errorf("error while getting url: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{
			URL:    url,
			Status: res.StatusCode,
		}
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading body: %w", err)
	}
	return string(body), nil
}
