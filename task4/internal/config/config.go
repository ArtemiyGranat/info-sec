package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Address string `yaml:"address"`
}

type Config struct {
	DbPath  string `yaml:"db-path"`
	Server         `yaml:"server"`
}

func ReadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
