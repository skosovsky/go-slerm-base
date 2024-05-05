package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload" // load .env
)

func signalHandlingHotReload(isDebug *string) {
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
				*isDebug = os.Getenv("MYAPP_DEBUG")
				log.Println("reread config, restart goroutines")
			}
		}
	}()

	<-done
}

func SetIsDebugEnv() {
	time.Sleep(10 * time.Second) //nolint:mnd // example
	err := os.Setenv("MYAPP_DEBUG", "1")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	log.Println(os.Getpid())

	isDebug := os.Getenv("MYAPP_DEBUG")

	go SetIsDebugEnv()

	signalHandlingHotReload(&isDebug)

	log.Println("isDebug:", isDebug)
}
