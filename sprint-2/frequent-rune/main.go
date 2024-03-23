package main

import "fmt"

func main() {
	data := "как дела?"
	fmt.Println(string(frequentRune(data)))
}

func frequentRune(str string) rune {
	frequentMap := make(map[rune]int)

	var maxCount int
	for _, v := range str {
		frequentMap[v]++

		if frequentMap[v] > maxCount {
			maxCount = frequentMap[v]
		}
	}

	var frequent rune
	for key, count := range frequentMap {
		if count == maxCount {
			frequent = key
		}
	}

	return frequent
}
