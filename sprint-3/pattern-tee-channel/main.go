package main

import (
	"context"
	"log"
	"time"
)

func tee(ctx context.Context, input <-chan string, outputs []chan<- string) {
	for elem := range input {
		for _, output := range outputs {
			go func() {
				select {
				case output <- elem:
					break
				case <-ctx.Done():
					break
				}
			}()
		}
	}
}

func main() {
	in := make(chan string)
	out1 := make(chan string)
	out2 := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	outputs := []chan<- string{out1, out2}

	go func() {
		in <- "A"
		in <- "B"
		in <- "C"
		close(in)
	}()

	tee(ctx, in, outputs)
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case elem := <-out1:
				log.Printf("out1 got value: %s\n", elem)
			case elem := <-out2:
				log.Printf("out2 got value: %s\n", elem)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
}
