package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = Client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
	}
}
