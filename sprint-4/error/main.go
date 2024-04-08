package main

import (
	"errors"
	"fmt"
	"log"
)

// type error interface {
// 	 Error() string
// }

func divide(a, b int) int {
	return a / b
}

func divide2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by 0")
	}

	return a / b, nil
}

var ErrDivisionZero = errors.New("division by 0")

func divide3(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionZero
	}

	return a / b, nil
}

type RetryError struct {
	NumRetries int
}

func NewRetryError(num int) RetryError {
	return RetryError{NumRetries: num}
}

func (r RetryError) Error() string {
	return fmt.Sprintf("Retries: %d", r.NumRetries)
}

type SomeOtherError struct {
	NumRetries int
}

func (s SomeOtherError) Error() string {
	return fmt.Sprintf("Other error with retries: %d", s.NumRetries)
}

func main() {
	// log.Println(divide(10,0)) // panic: runtime error, integer divide by zero

	val, err := divide2(10, 0)
	if err != nil {
		if err.Error() == "division by 0" {
			log.Println("Please provide valid input") // this
		} else {
			log.Printf("Some unknown error: %s\n", err)
		}
	} else {
		log.Printf("%d / %d = %d\n", 10, 0, val)
	}

	// sentinel error
	val, err = divide3(10, 0)
	if err != nil {
		if errors.Is(err, ErrDivisionZero) {
			log.Println("Please provide valid input") // this
		} else {
			log.Printf("Some unknown error: %s\n", err)
		}
	} else {
		log.Printf("%d / %d = %d\n", 10, 0, val)
	}

	// custom error
	err = NewRetryError(3)
	log.Println(err) // Retries: 3

	var retryErr RetryError
	if errors.As(err, &retryErr) {
		log.Println(retryErr) // Retries: 3
	}

	var otherErr SomeOtherError
	if errors.As(err, &otherErr) {
		log.Println(otherErr)
	}

	if e, ok := err.(RetryError); ok {
		log.Println("type assert err:", e) // type assert err: Retries: 3
	}
	if e, ok := err.(SomeOtherError); ok { // ok = false
		log.Println("type assert err:", e) // wouldn't execute
	}
}
