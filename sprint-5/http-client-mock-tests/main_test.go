package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestServiceGetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHTTPClient := NewMockHTTPClient(ctrl)

	service := &Service{
		HTTPClient: mockHTTPClient}

	expectedData := []byte("mocked data")

	t.Run("Success", func(t *testing.T) {
		expectedURL := "https://example.com/data"

		request, err := http.NewRequest(http.MethodGet, expectedURL, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		mockResponse := httptest.NewRecorder()
		_, err = mockResponse.Write(expectedData)
		if err != nil {
			t.Fatalf("failed to write response: %v", err)
		}

		mockHTTPClient.EXPECT().Do(request).Return(mockResponse.Result(), nil).Times(1)

		data, err := service.GetData(expectedURL)
		if err != nil {
			t.Fatalf("failed to get data: %v", err)
		}

		if !bytes.Equal(data, expectedData) {
			t.Errorf("expected %v, got %v", expectedData, data)
		}
	})

	t.Run("Error", func(t *testing.T) {
		expectedURL := "https://example.com/data"

		mockHTTPClient.EXPECT().Do(gomock.Any()).Return(nil, errors.New("HTTP client error")).Times(1)

		_, err := service.GetData(expectedURL)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		errHTTPClient := "could not send request: HTTP client error"
		if err.Error() != errHTTPClient {
			t.Errorf("expected %v, got %v", err, errHTTPClient)
		}
	})
}
