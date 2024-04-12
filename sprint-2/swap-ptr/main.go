package main

import (
	"log"
)

func main() {
	a := 5
	b := 10
	swap(&a, &b)
	log.Println(a, b)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
