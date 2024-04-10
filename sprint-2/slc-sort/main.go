package main

import (
	"fmt"
	"slices"
	"sort"
	"time"
)

func main() {
	slc := []int{2, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6, 1, 3, 7, 8, 4, 5, 6, 7, 8, 1, 9, 10, 11, 6}

	var start, finish time.Time

	fmt.Print("sortBubbleInPlace: ")
	start = time.Now()
	sortBubbleInPlace(slc)
	finish = time.Now()
	fmt.Println(finish.Sub(start))

	fmt.Print("sortBubbleReturn: ")
	start = time.Now()
	slc = sortBubbleReturn(slc)
	finish = time.Now()
	fmt.Println(finish.Sub(start))

	fmt.Print("sortReturn: ")
	start = time.Now()
	slc = sortReturn(slc)
	finish = time.Now()
	fmt.Println(finish.Sub(start))

	fmt.Print("sortStandard: ")
	start = time.Now()
	sortStandard(slc)
	finish = time.Now()
	fmt.Println(finish.Sub(start))

	fmt.Print("sortStandardParam: ")
	start = time.Now()
	sortStandardParam(slc)
	finish = time.Now()
	fmt.Println(finish.Sub(start))
}

func sortBubbleInPlace(slc []int) {
	if len(slc) <= 1 {
		return
	}

	n := len(slc)
	// Проходимся по всем элементам массива
	for i := 0; i < n-1; i++ {
		// Последний i элемент уже находится на своем месте
		for j := 0; j < n-i-1; j++ {
			// Сравниваем соседние элементы
			if slc[j] > slc[j+1] {
				// Меняем местами элементы, если они стоят в неправильном порядке
				slc[j], slc[j+1] = slc[j+1], slc[j]
			}
		}
	}
}

func sortBubbleReturn(slc []int) []int {
	if len(slc) <= 1 {
		return slc
	}

	n := len(slc)
	// Проходимся по всем элементам массива
	for i := 0; i < n-1; i++ {
		// Последний i элемент уже находится на своем месте
		for j := 0; j < n-i-1; j++ {
			// Сравниваем соседние элементы
			if slc[j] > slc[j+1] {
				// Меняем местами элементы, если они стоят в неправильном порядке
				slc[j], slc[j+1] = slc[j+1], slc[j]
			}
		}
	}

	return slc
}

func sortReturn(slc []int) []int {
	if len(slc) <= 1 {
		return slc
	}

	var sortedSlc []int

	minNum := slc[0]
	maxNum := slc[0]

	for i := range slc {
		if slc[i] < minNum {
			minNum = slc[i]
			continue
		}

		if slc[i] > maxNum {
			maxNum = slc[i]
		}
	}

	for i := minNum; i <= maxNum; i++ {
		countNum := 0
		for j := range slc {
			if slc[j] == i {
				countNum++
			}
		}

		for range countNum {
			sortedSlc = append(sortedSlc, i)
		}
	}

	return sortedSlc
}

func sortStandard(slc []int) {
	slices.Sort(slc)
}

func sortStandardParam(slc []int) {
	sort.SliceStable(slc, func(i, j int) bool {
		return slc[i] < slc[j]
	})
}
