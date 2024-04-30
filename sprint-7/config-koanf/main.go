package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload" // load .env
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func main() {
	k := koanf.New(".")

	if err := k.Load(file.Provider("example.json"), json.Parser()); err != nil {
		log.Fatal(err)
	}

	//TODO: Заменить везде на .env.example и заменить конфиг файлы
}
