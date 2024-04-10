package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		defer log.Println(scanner.Text())
	}
}
