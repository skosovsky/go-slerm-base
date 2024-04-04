package main

import (
	"fmt"
	"time"
)

type data struct {
	num int
	ptr *int
}

func main() {
	num10 := 10
	newData := data{num: 1, ptr: &num10}

	fmt.Println(newData.num, *newData.ptr) //nolint:forbidigo // it's learning code

	var ch = make(chan data)
	go readCh(ch)

	ch <- newData

	time.Sleep(time.Second)
	fmt.Println(newData.num, *newData.ptr) //nolint:forbidigo // it's learning code
}

func readCh(ch chan data) {
	num20 := 20
	res := <-ch

	res.num = 2
	*res.ptr = num20

	fmt.Println("read", res.num, *res.ptr) //nolint:forbidigo // it's learning code
}
