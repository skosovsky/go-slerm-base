package main

import (
	"log"
)

func main() {
	slc := []int{1, 2, 3}
	slc = appendNum(slc, 4, 5, 6) //nolint:mnd // it's learning code
	log.Print(slc)
	log.Println()

	slc4 := make([]int, 1, 4) //nolint:mnd // it's learning code
	changeSlc(slc4)
	log.Print(slc4)
}

func appendNum(slc []int, nums ...int) []int {
	return append(slc, nums...)
}

func changeSlc(slc []int) {
	log.Println(len(slc), cap(slc))
	slc[0] = 1
	log.Println(len(slc), cap(slc))
}
