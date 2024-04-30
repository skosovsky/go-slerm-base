package main

import (
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload" // load .env
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug   bool
	Port    int
	User    string
	Rate    float32
	Timeout time.Duration
	Users   []string
	Colors  map[string]int
}

func main() {
	// It's not work for app, because it for child
	// err := exec.Command("/bin/sh", "./config.sh").Run()
	// if err != nil {
	// 	 log.Fatal(err)
	// }

	var config Config
	err := envconfig.Process("myapp", &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(config.Debug, config.Port, config.User, config.Rate, config.Timeout)
	log.Println(config.Users)
	log.Println(config.Colors)
}
