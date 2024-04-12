package main

import (
	"log"
)

func main() {
	data := []string{"a", "b", "a"}
	log.Println(toFrequencyMap(data))
}

func toFrequencyMap(s []string) map[string]int {
	frequency := make(map[string]int)

	for _, v := range s {
		frequency[v]++
	}

	return frequency
}
