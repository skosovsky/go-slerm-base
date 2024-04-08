package main

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"log"
)

var Err1 = errors.New("error 1")
var Err2 = errors.New("error 2")

type CustomError struct{}

func (c CustomError) Error() string {
	return "custom error"
}

func step1() error {
	return Err1
}

func step2() error {
	return Err2
}

func step3() error {
	return &CustomError{}
}

func main() {
	var result error

	if err := step1(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step2(); err != nil {
		result = multierror.Append(result, err)
	}
	if err := step3(); err != nil {
		result = multierror.Append(result, err)
	}

	log.Println(result)
	// 3 errors occurred:
	//  * error 1
	//  * error 2
	//  * custom error

	if multiErr, ok := result.(*multierror.Error); ok {
		for _, e := range multiErr.Errors {
			if errors.Is(e, Err2) {
				log.Println("retry step2 may be?")
			}
		}
	}

	var customErr *CustomError
	if errors.As(result, &customErr) {
		log.Println(customErr) // custom error
	}

	if errors.Is(result, Err2) {
		log.Println("we have err 2 in list")
	}

	var multiErr *multierror.Error
	multiErr = multierror.Append(multiErr, Err1)
	log.Println(multiErr.ErrorOrNil())
}
