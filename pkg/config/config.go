package config

import (
	"log"

	"github.com/spf13/viper"
)

func NewConfig(filepath string) *viper.Viper {
	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigFile(filepath)

	err := config.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
	}

	return config
}
