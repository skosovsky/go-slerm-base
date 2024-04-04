package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"]) //nolint:forbidigo // it's learning code

	c.L.Unlock()
}

func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name) //nolint:forbidigo // it's learning code

	c.Broadcast()
	c.L.Unlock()
}

// See https://wcademy.ru/go-multithreading-sync-cond/
func main() {
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})

	go listen("слушатель 1", data, cond)
	go listen("слушатель 2", data, cond)

	go broadcast("источник", data, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
