package main

import "log"

// type iface struct {
// 	 tab  *itab
//	 data unsafe.Pointer
// }

type CustomError struct{}

func (c *CustomError) Error() string {
	return "custom error"
}

func nonNilErrorInterface() {
	var typed *CustomError = nil
	// var err error = (*CustomError)(nil)

	var err error = typed
	log.Println(typed == nil) // true
	if err != nil {
		log.Println(err) // custom err
	}
}

func main() {
	nonNilErrorInterface()
}
