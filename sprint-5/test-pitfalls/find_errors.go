package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func FindErrorsIn(file string) ([]string, error) {
	handle, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func(handle *os.File) {
		err = handle.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(handle)

	result := make([]string, 0)

	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "error") {
			result = append(result, text)
		}
	}
	return result, nil
}

func FindErrorsInWithIO(file string) ([]string, error) {
	handle, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func(handle *os.File) {
		err = handle.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(handle)

	return FindErrorsInWithoutIO(handle)
}

func FindErrorsInWithoutIO(reader io.Reader) ([]string, error) {
	result := make([]string, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "error") {
			result = append(result, text)
		}
	}
	return result, nil
}
