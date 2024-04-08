package main

import (
	"log"
	"os"
)

const filename = "/tmp/go-course.txt"

func invalidUsage() error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func correctUsage() error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func() {
		closeErr := f.Close()
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
