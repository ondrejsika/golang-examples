package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURI  = "mongodb://127.0.0.1:27017" // Replace with your MongoDB URI
	dbName    = "counter"
	collName  = "counter"
	counterID = "counter"
)

func main() {
	// MongoDB connection
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database(dbName).Collection(collName)

	// Query the current counter value
	filter := bson.M{"_id": counterID}
	update := bson.M{"$inc": bson.M{"count": 1}} // Increment the counter

	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)

	var result bson.M
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Print the updated counter value
	fmt.Printf("This script has been executed %v times.\n", result["count"])
}
