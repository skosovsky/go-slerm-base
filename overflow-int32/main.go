package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(add(1))
}

func add(a int32) int32 {
	return math.MaxInt32 + a
}
