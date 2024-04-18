package main

import (
	"log"
)

func main() {
	data := map[int]int{1: 1, 2: 2, 3: 3, 4: 4} //nolint:gomnd // it's learning code
	iterationByMap(data)
	log.Println(data)
}

func iterationByMap(data map[int]int) {
	for key, value := range data {
		log.Printf("%d:%d ", key, value)
	}
	log.Println()
}
