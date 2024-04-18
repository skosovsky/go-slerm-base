package main

import (
	"log"
)

func main() {
	myArr := [5]int{20, 15, 5, 30, 25}
	mySlice := myArr[1:4]

	log.Println(mySlice, len(mySlice), cap(mySlice)) // [15 5 30] 3 4

	myArrNew := [6]int{20, 15, 5, 30, 25, 40}
	mySliceNew := myArrNew[1:4]

	log.Println(mySliceNew, len(mySliceNew), cap(mySliceNew)) // [15 5 30] 3 5

	myArrNewest := [13]int{20, 15, 5, 30, 25, 40, 20, 7, 6, 5, 4, 1, 2}
	mySliceNewest := myArrNewest[1:4]

	log.Println(mySliceNewest, len(mySliceNewest), cap(mySliceNewest)) // [15 5 30] 3 12
}
