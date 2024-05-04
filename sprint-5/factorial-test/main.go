package main

import (
	"log"
)

func factorial(num int) int {
	result := 1

	for i := 1; i <= num; i++ {
		result *= i
	}

	return result
}

func main() {
	log.Println(factorial(10)) //nolint:mnd // it's learning code
}
