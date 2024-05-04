package main

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	_ "github.com/joho/godotenv/autoload" // load .env
	"github.com/spf13/viper"
)

func main() {
	vConfig := viper.New()
	vConfig.AddConfigPath("./")
	vConfig.SetConfigName("example")
	vConfig.SetConfigType("json")
	err := vConfig.ReadInConfig()
	if err != nil {
		log.Println(err)
	}

	appName := vConfig.Get("app_name")
	log.Println(appName)

	err = vConfig.WriteConfigAs("./config.yaml")
	if err != nil {
		log.Println(err)
	}

	vConfig.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	go vConfig.WatchConfig()

	envConfig := viper.New()
	envConfig.SetEnvPrefix("MYAPP")
	envConfig.AutomaticEnv()

	user := envConfig.GetString("USER")
	log.Println(user)

	// Don't work with env
	envConfig.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config env changed:", e.Name)
	})
	go vConfig.WatchConfig()

	time.Sleep(time.Second)
	err = os.Setenv("MYAPP_USER", "TEST")
	if err != nil {
		log.Println(err)
	}

	envConfig.AutomaticEnv()
	user = envConfig.GetString("USER")
	log.Println(user)

	select {} // for wait
}
