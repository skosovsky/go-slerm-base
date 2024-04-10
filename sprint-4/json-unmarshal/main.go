package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Example struct {
	A int    `json:"a"`
	B string `json:"B,omitempty"`
}

func jsonToStruct(jsonData []byte) (*Example, error) {
	structExample := Example{
		A: 0,
		B: "",
	}

	err := json.Unmarshal(jsonData, &structExample)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &structExample, nil
}

func main() {
	jsonData := `{"a":10,"B":"gold"}`
	structExample, err := jsonToStruct([]byte(jsonData))
	if err != nil {
		log.Println(err)
	}

	log.Println(*structExample)
}
