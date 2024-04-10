package main

import (
	"bufio"
	"io"
	"log"
)

type MyReader struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
		pos:  0,
	}
}

func (m *MyReader) Read(data []byte) (int, error) {
	if m.pos >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(data, m.data[m.pos:])
	m.pos += n

	return n, nil
}

func useScanner() {
	buffer := []byte("Hello, World!\nHappy New Year!")
	reader := NewReaderFromBuffer(buffer)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("error reading: %s", err)
	}
}

func customReader() {
	reader := NewReaderFromBuffer([]byte("Hello, World!\nHappy New Year!"))

	bufReader := bufio.NewReader(reader)

	for {
		line, err := bufReader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				log.Println(line)
			}
			break
		}
		if err != nil {
			log.Println("error reading:", err)
		}
		log.Print(line)
	}
}

func main() {
	useScanner()
	customReader()
}
