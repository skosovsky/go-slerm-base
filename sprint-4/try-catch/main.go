package main

import (
	"errors"
	"log"
	"time"
)

var ErrZeroDivisionAttempt = errors.New("divide by zero is not allowed")
var ErrNotAnError = errors.New("not an error")
var ErrEmptyArgument = errors.New("empty argument")

func divideForTry(a, b int) int {
	if b == 0 {
		panic(ErrZeroDivisionAttempt)
	}

	return a / b
}

func try(code func(), catch func(err error)) {
	defer func() {
		e := recover()
		if e != nil {
			err, ok := e.(error)
			if !ok {
				err = ErrNotAnError
			}
			catch(err)
		}
	}()

	code()
}

func divideForTryFunc(a, b int) {
	try(func() {
		result := divideForTry(a, b)
		log.Println(result)
	}, func(err error) {
		switch {
		case errors.Is(err, ErrZeroDivisionAttempt):
			// log.Println("Делить на ноль нельзя")
		case errors.Is(err, ErrEmptyArgument):
			// log.Println("Невозможно конвертировать пустую строку в число")
		}
	})
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrEmptyArgument
	}

	return a / b, nil
}

func divideFunc(a, b int) {
	result, err := divide(a, b)
	if err != nil {
		return
	}
	log.Println(result)
}

func main() {
	startTime := time.Now()
	for i := range 1000000 {
		divideForTryFunc(i, 0)
	}
	elapsedTime := time.Since(startTime)
	log.Println("Время выполнения с обработкой исключения:", elapsedTime)

	startTime = time.Now()
	for i := range 1000000 {
		divideFunc(i, 0)
	}
	elapsedTime = time.Since(startTime)
	log.Println("Время выполнения с обработкой ошибки:", elapsedTime)
}
