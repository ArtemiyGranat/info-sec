package storage

import (
	"context"
	"info-sec-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegistrateUser(db *mongo.Database, user models.User) error {
	collection := db.Collection(Users)

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
	    return err
	}

	return nil
}

func AuthUser(db *mongo.Database, username string) (*models.User, error) {
	collection := db.Collection(Users)
	filter := bson.M{"username": username}
	options := options.FindOne()
	
	var user models.User
	err := collection.FindOne(context.Background(), filter, options).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}