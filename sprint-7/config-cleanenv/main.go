package main

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload" // load .env
)

type Config struct {
	Debug   bool          `env:"MYAPP_DEBUG"   env-default:"false"`
	Port    int           `env:"MYAPP_PORT"    env-default:"8080"`
	User    string        `env:"MYAPP_USER"    env-default:"admin"`
	Rate    float32       `env:"MYAPP_RATE"    env-default:"0.0"`
	Timeout time.Duration `env:"MYAPP_TIMEOUT" env-default:"10s"`
	Users   []string      `env:"MYAPP_USERS"   env-default:"admin"`
}

type Conf struct {
	AppName  string `env:"APP_NAME" json:"app_name"`
	Debug    bool   `env:"DEBUG"    json:"debug"`
	Database struct {
		Host     string `env:"HOST"     json:"host"`
		Port     int    `env:"PORT"     json:"port"`
		User     string `env:"USER"     json:"user"`
		Password string `env:"password" json:"password"`
	} `env:"DATABASE" json:"database"`
}

func main() {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Println(err)
	}

	log.Println(config.User, config.Users, config.Debug, config.Port)

	var conf Conf
	err = cleanenv.ReadConfig("example.json", &conf)
	if err != nil {
		log.Println(err)
	}

	log.Println(conf.AppName, conf.Debug, conf.Database)
}
