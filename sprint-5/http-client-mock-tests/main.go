package main

import (
	"fmt"
	"io"
	"net/http"
)

//go:generate mockgen -source $GOFILE -destination ./mock.go -package ${GOPACKAGE}
// go generate ./...

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	HTTPClient
}

func (s *Service) GetData(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil) //nolint:noctx // it's learning code
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	return data, nil
}
