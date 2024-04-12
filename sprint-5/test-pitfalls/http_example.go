package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Value string `json:"value"`
}

type Result struct {
	Message string `json:"message"`
}

func sendRequest(request Request, url string) (*Result, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	body := bytes.NewBuffer(b)
	req, err := http.NewRequest(http.MethodPost, url, body) //nolint:noctx // it's learning code
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	b, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	defer response.Body.Close()

	result := Result{
		Message: "",
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &result, nil
}
