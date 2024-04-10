package main

import "log"

func divZero(zero int) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("it's recover from panic", err)
		}
	}()

	log.Println(10 / zero) //nolint:gomnd // it's learning code
}

func main() {
	divZero(0)

	log.Println("normal exit")
}
