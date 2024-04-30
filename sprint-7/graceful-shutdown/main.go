package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func gracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-sigs
		cancel()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if <-ctx.Done(); true {
			log.Println("Completing goroutine 1")
			time.Sleep(1 * time.Second)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if <-ctx.Done(); true {
			log.Println("Completing goroutine 2")
			time.Sleep(1 * time.Second)
			return
		}
	}()

	wg.Wait()
}

func main() {
	gracefulShutdown()
}
