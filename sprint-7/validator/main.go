package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	TimeZone string `env:"TIMEZONE" validate:"required"`

	AppMode   string `env:"APP_MODE"   validate:"required"`
	AppConfig string `env:"APP_CONFIG" validate:"required"`

	DBHost string `env:"DB_HOST" validate:"required"`
	DBPort int    `env:"DB_PORT" validate:"required,min=0,max=65535"`
	DBUser string `env:"DB_USER" validate:"required"`
	DBPass string `env:"DB_PASS" validate:"required"`
	DBName string `env:"DB_NAME" validate:"required"`

	ModemHost      string `env:"MODEM_HOST"       validate:"required,url"`
	ModemLoginUser string `env:"MODEM_LOGIN_USER" validate:"required"`
	ModemLoginPass string `env:"MODEM_LOGIN_PASS" validate:"required"`
}

func main() {
	cfg := Config{
		TimeZone:       "Local",
		AppMode:        "2",
		AppConfig:      "2",
		DBHost:         "2",
		DBPort:         1,
		DBUser:         "d",
		DBPass:         "d",
		DBName:         "d",
		ModemHost:      "ya.ru",
		ModemLoginUser: "d",
		ModemLoginPass: "d",
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(cfg)
	if err != nil {
		log.Println(err)
	}
}
