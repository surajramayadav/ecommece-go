package database

import (
	"context"
	"ecommerce/config"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func Connection() *mongo.Client {
	uri := config.MONGODB_URI

	if uri == "" {
		log.Fatal("Mongodb URI is empty")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	Client = client
	fmt.Println("================================connected with mongodb================================")

	// checking db already exists or not
	isDb := false

	listdb, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	fmt.Println("type checking", reflect.TypeOf(listdb))
	for _, name := range listdb {
		if name == "ecommerce" {
			isDb = true
		}

	}

	if isDb == false {
		fmt.Println("inside")
		// create a database
		db := client.Database("ecommerce")
		//create a collection
		db.CreateCollection(context.TODO(), "user")
		db.CreateCollection(context.TODO(), "product")
		db.CreateCollection(context.TODO(), "order")
	}

	return client

}
