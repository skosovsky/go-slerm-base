package main

import (
	"log"
	"strings"

	_ "github.com/joho/godotenv/autoload" // load .env
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func main() {
	config := koanf.New("/")

	if err := config.Load(file.Provider("example.json"), json.Parser()); err != nil {
		log.Println(err)
	}

	appName := config.String("app_name")
	debug := config.Bool("debug")
	dbHost := config.String("database.host")
	dbPort := config.Int("database.port")

	log.Println(appName, debug, dbHost, dbPort)

	err := config.Load(env.Provider("MYAPP_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "MYAPP_")), "_", ".")
	}), nil)
	if err != nil {
		log.Println(err)
	}

	users := config.String("users")
	log.Println(users)
}
