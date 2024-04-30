package main

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload" // load .env
)

type Config struct {
	Debug   bool           `env:"DEBUG"   envDefault:"false"`
	Port    int            `env:"PORT"    envDefault:"8080"`
	User    string         `env:"USER"    envDefault:"admin"`
	Rate    float32        `env:"RATE"    envDefault:"0.0"`
	Timeout time.Duration  `env:"TIMEOUT" envDefault:"10s"`
	Users   []string       `env:"USERS"   envDefault:"admin"`
	Colors  map[string]int `env:"COLORS"  envDefault:"0"`
}

func main() {
	var config Config
	opts := env.Options{
		Environment:           nil,
		TagName:               "",
		RequiredIfNoDef:       false,
		OnSet:                 nil,
		Prefix:                "MYAPP_",
		UseFieldNameByDefault: false,
		FuncMap:               nil,
	}

	if err := env.ParseWithOptions(&config, opts); err != nil {
		log.Println(err)
	}

	log.Println(config)

	cfg, err := env.ParseAsWithOptions[Config](opts) // it's generics
	if err != nil {
		log.Println(err)
	}

	log.Println(cfg)
}
