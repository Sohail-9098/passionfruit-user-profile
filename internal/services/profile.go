package services

import (
	"context"

	"github.com/Sohail-9098/passionfruit-user-profile/internal/config"
	"github.com/Sohail-9098/passionfruit-user-profile/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProfile(profile models.Profile, client *mongo.Client) error {
	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	if _, err := collection.InsertOne(context.TODO(), profile); err != nil {
		return err
	}
	return nil
}
