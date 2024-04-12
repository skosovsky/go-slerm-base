package main

import (
	"log"
	"strings"
	"unicode"
)

func step1(in <-chan string, out chan<- string) {
	go func() {
		defer close(out)
		for value := range in {
			words := strings.Fields(value)
			out <- strings.Join(words, " ")
		}
	}()
}

func step2(in <-chan string, out chan<- string) {
	go func() {
		defer close(out)
		for value := range in {
			lines := strings.Split(value, ".")
			for _, line := range lines {
				if len(line) == 0 {
					continue
				}

				out <- strings.Trim(line, " ")
			}
		}
	}()
}

func step3(in <-chan string) <-chan string {
	var out = make(chan string)
	go func() {
		defer close(out)
		for value := range in {
			if len(value) == 0 {
				continue
			}
			runes := []rune(value)
			firstLetter := unicode.ToUpper(runes[0])
			out <- string(firstLetter) + string(runes[1:])
		}
	}()

	return out
}

func generateWork(works []string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for _, work := range works {
			ch <- work
		}
	}()

	return ch
}

func main() {
	works := []string{" Привет, как дела.  Что нового у тебя происходит.    И где ты живешь  ",
		" Что нового у тебя происходит. Привет, как дела.  где ты живешь. Что нового у тебя происходит.    И где ты живешь  ",
		" Что нового у тебя происходит. Привет, как дела.  Что нового у тебя происходит.    И где ты живешь  "}

	pipeIn := generateWork(works)

	var out1 = make(chan string)
	step1(pipeIn, out1)

	var out2 = make(chan string)
	step2(out1, out2)

	out3 := step3(out2)
	for v := range out3 {
		log.Println(v)
	}
}
