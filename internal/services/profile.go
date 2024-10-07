package services

import (
	"context"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/config"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProfile(client *mongo.Client, profile models.Profile) (string, error) {
	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	result, err := collection.InsertOne(context.TODO(), profile)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func UpdateProfile(client *mongo.Client, profile models.Profile, id string) error {
	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	if _, err := collection.ReplaceOne(context.TODO(), filter, profile); err != nil {
		return err
	}
	return nil
}
