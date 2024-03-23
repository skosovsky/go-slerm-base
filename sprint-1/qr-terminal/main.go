package main

import (
	"github.com/mdp/qrterminal/v3"
	"os"
)

func main() {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig("Hello, world", config)
}

// go build - компиляция с использованием file для вывода информации
// GOOS=linux GOARCH=amd64 go build -o qr-terminal-linux - под linux
// file qr-terminal-linux - выводим информацию о файле
// go build -ldflags="-s -w" -o qr-terminal - компилируем без отладочной информации
