package tools

import (
	"go/types"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *types.Config {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Fatal error config file: %v ", err)
	}
	var config types.Config
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
