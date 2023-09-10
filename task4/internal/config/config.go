package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Address string `yaml:"address"`
}

type Config struct {
	Server `yaml:"http_server"`
}

func ReadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Could not read the config file %s: %v", configPath, err.Error())
	}

	return &cfg, nil
}
