package main

import (
	"fmt"
	"log"
)

func main() {
	slc := []int{1, 2, 3}
	slc = appendNum(slc, 4, 5, 6)
	log.Print(slc)
	log.Println()

	slc4 := make([]int, 1, 4)
	changeSlc(slc4)
	fmt.Print(slc4)
}

func appendNum(slc []int, nums ...int) []int {
	return append(slc, nums...)
}

func changeSlc(slc []int) {
	log.Println(len(slc), cap(slc))
	slc[0] = 1
	log.Println(len(slc), cap(slc))
}
