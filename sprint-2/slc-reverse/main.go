package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(reverse(slc))
}

func reverse(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	var sliceReversed []int

	for i := len(slice) - 1; i >= 0; i-- {
		sliceReversed = append(sliceReversed, slice[i])
	}

	return sliceReversed
}
