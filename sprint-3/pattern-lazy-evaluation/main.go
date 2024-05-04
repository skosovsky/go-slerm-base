package main

import (
	"log"
	"sync"
)

type LazyInt func() int

func Make(numInt LazyInt) LazyInt {
	var num int
	var once sync.Once

	return func() int {
		once.Do(func() {
			num = numInt()
			numInt = nil // so that f can now be GC-ed
		})
		return num
	}
}

func main() {
	lazyInt := Make(func() int {
		log.Println("Doing expensive calculations")
		return 23 //nolint:mnd // it's learning code
	})

	log.Println(lazyInt())
	log.Println(lazyInt() + 42) //nolint:mnd // it's learning code
}
