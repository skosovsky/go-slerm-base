package main

import "log"

type Data struct {
	value int
}

func (d Data) PrintValue() {
	log.Println(d.value)
}

func (d *Data) PrintValuePtr() {
	log.Println(d.value)
}

func deferMethods() {
	c := Data{value: 123} //nolint:mnd // it's learning code
	defer c.PrintValue()  // 123
	c.value = 456

	c2 := Data{value: 123}   //nolint:mnd // it's learning code
	defer c2.PrintValuePtr() // 456
	c2.value = 456
}

func main() {
	deferMethods()
}
