package main

import (
	"log"
)

func main() {
	data1 := map[int]struct{}{1: {}, 2: {}, 21: {}}
	data2 := map[int]struct{}{21: {}}
	intersect := mapKeyIntersect(data1, data2)
	log.Println(intersect)
}

func mapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	var data []int
	for i := range m1 {
		_, ok := m2[i]
		if ok {
			data = append(data, i)
		}
	}

	return data
}
