package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(url string) (*mongo.Client, error) {
	// create a database connection
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := client.Connect(context.TODO()); err != nil {
		return nil, err
	}

	return client, nil
}
