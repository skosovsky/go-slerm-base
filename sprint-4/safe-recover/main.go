package main

import (
	"log"
	"time"
)

type Request struct {
	Payload string
}

func server(c <-chan Request) {
	for work := range c {
		go safelyDo(work)
	}
}

func safelyDo(req Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("request error:", err)
		}
	}()
	do(req)
}

func do(r Request) {
	log.Println("handle request:", r.Payload)
	if r.Payload == "do-panic" {
		panic("failed to process this request")
	}
	log.Println("done handling request:", r.Payload)
}

func safeRecover() {
	c := make(chan Request)
	go server(c)
	c <- Request{Payload: "success 1"}
	c <- Request{Payload: "do-panic"}
	c <- Request{Payload: "success 2"}
	defer close(c)
}

// handle request: success 1
// done handling request: success 1

// handle request: success 2
// done handling request: success 2

// handle request: do-panic
// request error: failed to process this request

func main() {
	safeRecover()
	time.Sleep(time.Second)
}
