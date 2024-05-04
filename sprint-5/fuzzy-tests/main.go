package main

import (
	"log"

	"github.com/skosovsky/go-slerm-base/sprint-5/fuzzy-tests/math"
)

func main() {
	result := math.AddWithError(2, 2)   //nolint:mnd // it's learning code
	log.Println(result)                 // 4
	result = math.AddWithError(100, 10) //nolint:mnd // it's learning code
	log.Println(result)                 // 0
}
