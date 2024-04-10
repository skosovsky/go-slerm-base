package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

func reverseReaderDefer(reader io.Reader) (result []string, err error) { //nolint:nonamedreturns // it's defer
	var data []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	for _, v := range data {
		defer func(result *[]string) {
			*result = append(*result, v)
		}(&result)
	}

	return result, nil
}

func main() {
	str := []byte("1\n2\n3\n4\n5\n6")
	reader := bytes.NewReader(str)

	linesReverse, err := reverseReaderDefer(reader)
	if err != nil {
		log.Println(err)
	}

	log.Println(linesReverse)
}
