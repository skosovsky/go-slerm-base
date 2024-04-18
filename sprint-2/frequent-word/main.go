package main

import (
	"log"
	"strings"
)

func main() {
	str := "привет мой дорогой друг и снова привет мой лучший друг но ты же друг"
	log.Println(frequentWord(str))
}

func frequentWord(str string) string {
	sentence := strings.Split(str, " ")

	words := make(map[string]int)
	for _, word := range sentence {
		words[word]++
	}

	var mostFrequentWord string
	var maxCount int
	for word, count := range words {
		if count > maxCount {
			maxCount = count
			mostFrequentWord = word
		}
	}

	return mostFrequentWord
}
