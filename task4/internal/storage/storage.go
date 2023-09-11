package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DbName = "info-sec-api"
	Users  = "users" 
)

func Connect(storagePath string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(storagePath)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(DbName)

	// Only for debug, remove it later
	// _, err = db.Collection(Users).DeleteMany(context.Background(), bson.M{})
	return db, nil
}

func Close(db *mongo.Database) error {
	if err := db.Client().Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}