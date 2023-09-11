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
		return err
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	db, err := storage.Connect(cfg.DbPath)
	if err != nil {
		return err
	}
	defer storage.Close(db)

	router := routes.SetupRouter(db, cfg.Address)
	err = router.Run(cfg.Address + ":" + cfg.Port)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("Could not run the server: %v\n", err)
	}
}