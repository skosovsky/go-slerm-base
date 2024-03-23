package main

import "fmt"

func main() {
	data := "Привет, как дела?"
	fmt.Println(removeSpaces(data))
}

func removeSpaces(s string) string {
	sRune := []rune(s)
	var sRuneNoSpaces []rune

	for i := range sRune {
		if sRune[i] != ' ' {
			sRuneNoSpaces = append(sRuneNoSpaces, sRune[i])
		}
	}

	return string(sRuneNoSpaces)
}
