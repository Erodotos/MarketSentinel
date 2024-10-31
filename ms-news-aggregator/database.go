package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mc *MongoInstance

var MONGO_HOST = os.Getenv("MONGO_HOST")
var MONGO_DB = os.Getenv("MONGO_DB")
var MONGO_COLLECTION = os.Getenv("MONGO_COLLECTION")

func MongoClient() (MongoInstance, error) {

	// Reuse Connection if exists
	if mc != nil {
		return *mc, nil
	}

	// Create new Mongo instance
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_HOST))
	if err != nil {
		return MongoInstance{}, err
	}

	// Check database connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error().Msg(fmt.Sprintf("Couldn't connect to Mongo: %s/%s", MONGO_HOST, MONGO_DB))
	}
	logger.Info().Msg(fmt.Sprintf("Connected  to Mongo: %s/%s", MONGO_HOST, MONGO_DB))

	// Create Mongo Instance
	mc = &MongoInstance{
		Client: client,
		Db:     client.Database(MONGO_DB),
	}

	return *mc, nil
}
