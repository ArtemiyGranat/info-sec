package main

import (
	"info-sec-api/internal/config"
	"info-sec-api/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func run() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load .env file: %v", err)
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	router := routes.SetupRouter()
	err = router.Run(cfg.Address)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Could not run the server: %v\n", err)
	}

	return
}