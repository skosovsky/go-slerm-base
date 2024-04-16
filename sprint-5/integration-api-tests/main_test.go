package main

import (
	"net/http"
	"testing"
)

func TestClientAPI_GetData(t *testing.T) {
	resp, err := GetData("https://example.com/data")
	if err != nil {
		t.Fatalf("Error getting data: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("GetData returned wrong status code: got %v want %v", resp.StatusCode, http.StatusNotFound)
	}
}
