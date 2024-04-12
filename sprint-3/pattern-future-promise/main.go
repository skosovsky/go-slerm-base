package main

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

type data struct {
	Body string
	Err  error
}

func doGet(url string) (string, error) {
	time.Sleep(200 * time.Millisecond) //nolint:gomnd // it's learning code

	failure := rand.Int()%10 > 5 //nolint:gosec // it's learning code
	if failure {
		return "", errors.New("timeout")
	}

	return "Response of " + url, nil
}

func future(url string) <-chan data {
	ch := make(chan data, 1)

	go func() {
		for range 3 {
			body, err := doGet(url)
			if err != nil {
				log.Println("got error", err, "retrying")
				time.Sleep(10 * time.Millisecond) //nolint:gomnd // it's learning code
				continue
			}
			ch <- data{Body: body, Err: err}
		}
	}()

	return ch
}

func main() {
	future1 := future("https://example.com")
	future2 := future("https://google.com")

	log.Println("Requests started")

	body1 := <-future1
	body2 := <-future2

	log.Printf("Response 1: %v\n", body1)
	log.Printf("Response 2: %v\n", body2)
}
