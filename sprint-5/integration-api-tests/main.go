package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetData(address string) (*http.Response, error) {
	resp, err := http.Get(address) //nolint:gosec,noctx // it's learning code
	if err != nil {
		return &http.Response{}, fmt.Errorf("error fetching data from %s: %w", address, err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}(resp.Body)

	return resp, nil
}

func main() {

}
