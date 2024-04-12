package main

import (
	"log"
	"math"
)

func main() {
	log.Println(math.MaxInt32)
	log.Println(add(1))
}

func add(a int32) int32 {
	return math.MaxInt32 + a
}
