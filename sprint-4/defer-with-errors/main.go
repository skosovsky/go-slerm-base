package main

import (
	"fmt"
	"log"
	"os"
)

const filename = "/tmp/go-course.txt"

func invalidUsage() error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening file error: %w", err)
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	return nil
}

func correctUsage() error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("opening file error: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			if err == nil {
				err = closeErr
			} else {
				log.Println("error occurred closing the file:", closeErr)
			}
		}
	}()

	return err
}

func deferWithErrors() {
	log.Println(invalidUsage())
	log.Println(correctUsage())
}

func main() {
	deferWithErrors()
}
