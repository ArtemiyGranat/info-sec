package main

import (
	"info-sec-api/internal/config"
	"info-sec-api/internal/routes"
	"info-sec-api/internal/storage"
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

	db, err := storage.Connect(cfg.DbPath)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer storage.Close(db)

	router := routes.SetupRouter(db)
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
}