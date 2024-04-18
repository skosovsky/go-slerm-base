package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Value string `json:"value"`
}

type Result struct {
	Message string `json:"message"`
}

func sendRequest(request Request, url string) (*Result, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest(http.MethodPost, url, body) //nolint:noctx // it's learning code
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req) //nolint:bodyclose // body is closed
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	data, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}(response.Body)

	result := Result{
		Message: "",
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}
