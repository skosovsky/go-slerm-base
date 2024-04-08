package main

import (
	"log"
	"runtime"
)

func main() {
	log.Print(runtime.GOMAXPROCS(0)) // 10
}
