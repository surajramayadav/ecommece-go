package database

import (
	"context"
	"fmt"
	"instant/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client {
	uri := config.MONGODB_URI

	if uri == "" {
		log.Fatal("Mongodb URI is empty")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("================================connected with mongodb================================")

	return client

}

// Client Database instance
var Client *mongo.Client = Connection()

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("instant").Collection(collectionName)

	return collection
}
