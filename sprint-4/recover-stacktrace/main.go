package main

import (
	"log"
	"runtime/debug"
)

func recoverFromPanic() {
	err := recover()
	if err != nil {
		log.Println("it's recover from panic", err)
		log.Println("Stack trace:", string(debug.Stack()))
	}
}

func divZero(zero int) {
	defer recoverFromPanic()
	log.Println(10 / zero) //nolint:gomnd // it's learning code
}

func main() {
	divZero(0)

	log.Println("normal exit")
}
