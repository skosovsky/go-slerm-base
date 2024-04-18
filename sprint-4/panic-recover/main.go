package main

import (
	"log"
	"math/rand"
)

func unrecoverableError() {
	panic("stop")
}

func handePanic() {
	err := recover()
	if err != nil {
		log.Printf("recovered from panic: %s\n", err)
	}
}

func main() {
	var i = rand.Int() % 10 //nolint:gosec // it's learning code
	var a [8]int

	// panic: runtime error: index out of range [9] with length 8
	log.Println(a[i])

	defer log.Println(1)
	defer log.Println(2) //nolint:gomnd // it's learning code
	defer log.Println(3) //nolint:gomnd // it's learning code

	defer handePanic() // recovered from error stop
	unrecoverableError()

	// 0
	// recover from error stop
	// 3
	// 2
	// 1
}
