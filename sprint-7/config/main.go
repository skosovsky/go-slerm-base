package main

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
)

func yamlConfigExample() {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
server:
  addr: "127.0.0.1"
  port: 8080
middlewares:
  - "csrf"
  - "rate-limiter"
apiKey: "secret"
`)

	type yamlConfig struct {
		Server struct {
			Addr string `yaml:"addr"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
		Middlewares []string `yaml:"middlewares"`
		APIKey      string   `yaml:"secret"`
	}

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		log.Panic(err)
	}

	config := yamlConfig{} //nolint:exhaustruct // it's new empty object
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Panic(err)
	}
	log.Println(config) // 2024/04/26 07:18:24 {{127.0.0.1 8080} [csrf rate-limiter] secret}
}

func configFromEnv() {
	v := viper.New() //nolint:varnamelen // it's learning code

	type envConfig struct {
		ServerAddr  string   `mapstructure:"SERVER_ADDR"`
		ServerPort  int      `mapstructure:"SERVER_PORT"`
		Middlewares []string `mapstructure:"MIDDLEWARES"`
		APIKey      string   `mapstructure:"SECRET"`
	}

	// don't confuse
	// viper.SetConfigType("env")

	v.SetConfigType("env")
	v.SetEnvPrefix("MY_SERVICE")

	v.SetDefault("SECRET", "")
	v.SetDefault("MIDDLEWARES", []string{})
	v.SetDefault("SERVER_ADDR", "127.0.0.1")
	v.SetDefault("SERVER_PORT", 9090) //nolint:mnd // it's example

	v.AutomaticEnv()
	cfg := envConfig{} //nolint:exhaustruct // it's new empty object
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	// export MY_SERVICE_SERVER_ADDR="10.0.0.1"
	// export MY_SERVICE_SERVER_PORT="9090"
	// export MY_SERVICE_SECRET=env-secret
	log.Println(cfg) // 2024/04/26 08:18:17 {10.0.0.1 9090 [] env-secret}
}

func configFromEnvAndFromFile() {
	v := viper.New() //nolint:varnamelen // it's learning code

	type envAndFileConfig struct {
		ServerAddr  string   `mapstructure:"SERVER_ADDR" yaml:"server_addr"`
		ServerPort  int      `mapstructure:"SERVER_PORT" yaml:"server_port"`
		Middlewares []string `mapstructure:"MIDDLEWARES" yaml:"middlewares"`
		APIKey      string   `mapstructure:"SECRET"      yaml:"secret"`
	}

	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	v.AutomaticEnv()
	cfg := envAndFileConfig{} //nolint:exhaustruct // it's new empty object
	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	log.Println(cfg) // 2024/04/26 08:22:18 {127.0.0.1 8080 [csrf ip-blacklist rate-limiter] config-secret}

	// export SECRET=other
	log.Println(cfg) // 2024/04/26 08:24:10 {127.0.0.1 8080 [csrf ip-blacklist rate-limiter] other}
}

func main() {
	yamlConfigExample()

	configFromEnv()

	configFromEnvAndFromFile()
}
