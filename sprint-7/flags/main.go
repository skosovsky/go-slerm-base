package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

var listenPortFlag = flag.Int("port", 8080, "port to listen on") //nolint:gochecknoglobals // it's learning code
var listenAndFlag string                                         //nolint:gochecknoglobals // it's learning code

func readEnvSimple() {
	listerPortStr := os.Getenv("PORT")
	if listerPortStr == "" {
		log.Panic("No PORT variable set")
	}
	listerPort, err := strconv.Atoi(listerPortStr)
	if err != nil {
		log.Panic(err)
	}
	log.Println(listerPort)

	err = os.Setenv("PORT", "80") // bad practice
	if err != nil {
		log.Panic(err)
	}
}

func useFlagsSimple() {
	// go run . --help
	// Usage of app:
	// -addr string
	// Addr to listen (default "127.0.0.1")
	// -port int
	// port to listen on (default 8080)

	// app -addr 10.10.10.10
	flag.StringVar(&listenAndFlag, "addr", "127.0.0.1", "Addr to listen")
	log.Println(listenAndFlag, *listenPortFlag)
}

func main() {
	flag.Parse()

	readEnvSimple()

	useFlagsSimple()
}
