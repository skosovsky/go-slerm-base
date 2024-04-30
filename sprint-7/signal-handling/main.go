package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func signalHandling() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	// kill -TERM 97716
	go func() {
		sig := <-sigs
		if sig == syscall.SIGTERM {
			log.Println("Received SIGTERM, shutting down", sig) // in this place need close connection and cancel context
			done <- true
		}
	}()

	log.Println("Waiting for SIGINT/SIGTERM to exit")
	<-done
	log.Println("Received SIGINT/SIGTERM, shutting down")
}

func signalHandlingAsNginx() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGTERM)

	done := make(chan bool, 1)

	// kill -HUP 97943 // Получив сигнал HUP nginx перечитывает конфиг, можно использовать в своем приложении
	go func() {
		for {
			sig := <-sigs
			if sig == syscall.SIGTERM {
				close(done)
				return
			}
			if sig == syscall.SIGHUP {
				log.Println("reread config, restart goroutines")
			}
		}
	}()

	<-done
}

func main() {
	signalHandling()
	signalHandlingAsNginx()
}
