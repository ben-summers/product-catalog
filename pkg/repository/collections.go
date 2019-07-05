package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"product-catalog/pkg/config"
)

func getConfiguration() *config.Configuration {
	return config.Settings
}

func getDatabaseConfiguration() *config.Database {
	return &getConfiguration().Database
}

func getCollection(collectionName string) *mongo.Collection {
	cfg := getConfiguration()
	return MongoClient.Database(cfg.Database.DatabaseName).Collection(collectionName)
}

func getProductsCollection() *mongo.Collection {
	return getCollection(getDatabaseConfiguration().ProductCollectionName)
}
