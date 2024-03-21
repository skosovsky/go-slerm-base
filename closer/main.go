package main

import "fmt"

func main() {
	a := funcWithCloser()
	fmt.Println(a) // 0x1026c8bd0 // fmt.Println arg a is a func value, not called

	fmt.Println(a(10)) // 10
	fmt.Println(a(20)) // 10 + 20 = 30
}

func funcWithCloser() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v

		return sum
	}
}
