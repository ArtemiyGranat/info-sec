package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "info-sec-api"
)

func Connect(storagePath string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(storagePath)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(DbName)
	return db, nil
}

func Close(db *mongo.Database) error {
	if err := db.Client().Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}