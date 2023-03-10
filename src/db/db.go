package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Database = nil

func ConnectDB() *mongo.Database {

	if client != nil {
		return client
	}

	var MONGO_URL string = os.Getenv("MONGO_URL")

	log.Println("Connecting to MongoDB...")

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(MONGO_URL).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to MongoDB!")
	}

	return client.Database("beaurl")
}

func GetInstance() *mongo.Database {
	if client == nil {
		client = ConnectDB()
	}

	return client
}
