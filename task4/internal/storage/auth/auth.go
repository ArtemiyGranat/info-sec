package storage

import (
	"context"
	"info-sec-api/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func RegistrateUser(db *mongo.Database, user models.User) error {
	collection := db.Collection("users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
	    return err
	}

	return nil
}