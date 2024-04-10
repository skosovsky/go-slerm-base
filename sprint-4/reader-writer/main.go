package main

import (
	"errors"
	"io"
	"log"
)

type MyReader struct {
	data []byte
	pos  int
}

type MyWriter struct {
	data []byte
	pos  int
}

func NewReaderFromBuffer(buffer []byte) *MyReader {
	return &MyReader{
		data: buffer,
		pos:  0,
	}
}

func NewWriterToBuffer(buffer []byte) *MyWriter {
	return &MyWriter{
		data: buffer,
		pos:  0,
	}
}

func (m *MyReader) Read(p []byte) (int, error) {
	if m.pos >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(p, m.data[m.pos:])
	m.pos += n

	return n, nil
}

func (m *MyWriter) Write(p []byte) (int, error) {
	if m.pos+len(p) < len(m.data) {
		return 0, errors.New("buffer overflow")
	}

	n := copy(m.data[m.pos:], p)
	m.pos += n

	return n, nil
}

func main() {
	reader := NewReaderFromBuffer([]byte("Hello, World!\nHappy New Year!"))
	buffer := make([]byte, len(reader.data))
	writer := NewWriterToBuffer(buffer)

	n, err := io.Copy(writer, reader)
	if err != nil {
		log.Println("copy error:", err)
	}

	log.Println(n)
}
