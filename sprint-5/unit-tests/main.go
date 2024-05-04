package main

import (
	"log"

	"github.com/skosovsky/go-slerm-base/sprint-5/unit-tests/math"
)

func main() {
	result := math.Add(2, 2) //nolint:mnd // it's learning code
	log.Println(result)
}
