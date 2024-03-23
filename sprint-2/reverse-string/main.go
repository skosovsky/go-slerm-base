package main

import "fmt"

func main() {
	data := "привет!"
	fmt.Println(reverse(data))
}

func reverse(s string) string {
	sRune := []rune(s)
	var sReverse []rune

	for i := len(sRune) - 1; i >= 0; i-- {
		sReverse = append(sReverse, sRune[i])
	}

	return string(sReverse)
}
