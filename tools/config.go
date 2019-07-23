package tools

import (
	"log"

	"github.com/spf13/viper"
)

type Server struct {
	Url string
}

type Config struct {
	Server Server
}

func LoadConfig() *Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: %v ", err)
	}
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
