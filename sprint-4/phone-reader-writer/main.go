package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

const phoneLength = 11

type phoneReader string

func (r phoneReader) Read(data []byte) (int, error) {
	count := 0
	for i := range r {
		if r[i] >= '0' && r[i] <= '9' {
			if count >= phoneLength {
				return count, errors.New("invalid phone number")
			}

			data[count] = r[i]
			count++
		}
	}

	return count, io.EOF
}

type phoneWriter struct{}

func (w phoneWriter) Write(data []byte) (int, error) {
	if len(data) == 0 {
		return 0, nil
	}

	for i := range data {
		if data[i] >= '0' && data[i] <= '9' {
			fmt.Print(string(data[i])) //nolint:forbidigo // it's learning code
		}
	}
	fmt.Println() //nolint:forbidigo // it's learning code
	return len(data), nil
}

func main() {
	phone1 := phoneReader("+7(964)218 93310")

	buffer := make([]byte, phoneLength)
	n, err := phone1.Read(buffer)

	log.Println(string(buffer))

	log.Println(n, err, len(buffer), cap(buffer))

	phone2 := []byte("+7(964)218 9310")
	writer := phoneWriter{}
	n, err = writer.Write(phone2)

	log.Println(string(phone2))

	log.Println(n, err, len(phone2), cap(phone2))
}
