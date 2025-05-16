package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TCP struct {
		Protocol string
		Address  string
		Port     string
	}
	Emulator struct {
		TimeInterval                  int
		NumberOfSources               int
		NumberOfInstancesOfEachSource int
	}
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %s", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode config: %s", err)
		return nil, err
	}

	return &config, nil
}
