package main

import "fmt"

func main() {
	a := 5
	b := 10
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
