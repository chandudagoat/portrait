package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func GetCollection(collection string) *mongo.Collection {
	return MongoClient.Database("Cluster0").Collection(collection)
}

func ConnectMongo() error {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("Could not load MONGO_URI from .env file")
	}

	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return nil
}

func DisconnectMongo() {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
