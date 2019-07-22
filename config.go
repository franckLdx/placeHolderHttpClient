package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
}

func LoadConfig() *Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: %s ", err)
	}
	var config Config
	v.Unmarshal(&config)
	return &config
}
