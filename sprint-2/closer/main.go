package main

import (
	"log"
)

func main() {
	a := funcWithCloser()
	// log.Println(a) // 0x1026c8bd0 // fmt.Println arg a is a func value, not called

	log.Println(a(10)) //nolint:mnd // 10
	log.Println(a(20)) //nolint:mnd // 10 + 20 = 30
}

func funcWithCloser() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}
