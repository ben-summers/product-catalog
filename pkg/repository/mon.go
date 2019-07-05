package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	MongoClient *mongo.Client
)

func ConfigureMongo(connectionString string) error {
	if MongoClient != nil {
		return errors.New("mongo client already configured")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	MongoClient = client

	err = MongoClient.Ping(context.TODO(), nil)
	if err == nil {
		log.Printf("database host at %s is reachable.", connectionString)
	}

	return nil
}
