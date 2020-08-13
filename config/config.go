package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Port int
}

func GetConfig() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return conf, err
	}
	return conf, nil
}

func NewConfig() (Config, error) {
	return GetConfig()
}
