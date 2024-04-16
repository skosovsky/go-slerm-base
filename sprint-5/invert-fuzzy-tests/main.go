package main

import (
	"errors"
	"log"
	"unicode/utf8"
)

func main() {
	log.Println(invertLine("hello"))
	log.Println(invertLine("привет"))
	log.Println(invertLine("\xff"))
}

func invertLine(line string) (string, error) {
	if !utf8.ValidString(line) {
		return line, errors.New("not a valid UTF-8 string")
	}
	runes := []rune(line)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}
