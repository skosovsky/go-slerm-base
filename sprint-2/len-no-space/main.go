package main

import "fmt"

func main() {
	data := "Привет, как дела?"
	fmt.Println(stringLengthWithoutSpaces(data))
}

func stringLengthWithoutSpaces(str string) int {
	sRune := []rune(str)
	var sRuneNoSpaces []rune

	for i := range sRune {
		if sRune[i] != ' ' && sRune[i] != '	' {
			sRuneNoSpaces = append(sRuneNoSpaces, sRune[i])
		}
	}

	return len(sRuneNoSpaces)
}
