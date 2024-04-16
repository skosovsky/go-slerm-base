package main

import (
	"log"
)

func main() {
	slc := []int{0, 0, 1, 2, 2, 1, 3, 4, 5, 6, 6, 6, 7, 6}
	log.Println(removeDuplicates(slc))
}

func removeDuplicates(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	var sliceWithoutDuplicates []int

	for i := range slice {
		if len(sliceWithoutDuplicates) == 0 {
			sliceWithoutDuplicates = append(sliceWithoutDuplicates, slice[i])
			continue
		}

		var countFound int
		for j := range sliceWithoutDuplicates {
			if slice[i] == sliceWithoutDuplicates[j] {
				countFound++
			}
		}

		if countFound == 0 {
			sliceWithoutDuplicates = append(sliceWithoutDuplicates, slice[i])
		}
	}

	return sliceWithoutDuplicates
}
