package main

import (
	"fmt"
	"net/http"
)

func GetData(address string) (*http.Response, error) {
	resp, err := http.Get(address)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error fetching data from %s: %w", address, err)
	}
	defer resp.Body.Close()

	return resp, nil
}

func main() {

}
