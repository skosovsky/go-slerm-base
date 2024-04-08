package main

import (
	"log"
	"sync"
)

type LazyInt func() int

func Make(f LazyInt) LazyInt {
	var num int
	var once sync.Once

	return func() int {
		once.Do(func() {
			num = f()
			f = nil // so that f can now be GC-ed
		})
		return num
	}
}

func main() {
	n := Make(func() int {
		log.Println("Doing expensive calculations")
		return 23
	})

	log.Println(n())
	log.Println(n() + 42)
}
