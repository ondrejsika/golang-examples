package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/rand"
)

func main() {
	rand.Seed(uint64(time.Now().UnixNano()))

	// MongoDB connection URI (using 127.0.0.1)
	uri := "mongodb://127.0.0.1:27017" // MongoDB running locally

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// Select the database and collection
	database := client.Database("mydb")
	collection := database.Collection("mycollection")

	// Create a document to insert
	doc := bson.D{
		{Key: "name", Value: getRandomDogName()},
		{Key: "age", Value: getRandomDogAge()},
		{Key: "breed", Value: getRandomDogBreed()},
	}

	// Insert the document into the collection
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	// Find all documents in the collection
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	fmt.Println("All documents in the collection:")
	for cursor.Next(ctx) {
		var resultDoc bson.M
		if err := cursor.Decode(&resultDoc); err != nil {
			log.Fatal(err)
		}
		fmt.Println(
			resultDoc["_id"],
			resultDoc["name"],
			resultDoc["age"],
			resultDoc["breed"],
		)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
}

func getRandomDogName() string {
	names := []string{"Buddy", "Charlie", "Daisy", "Lucy", "Max", "Molly", "Rocky"}
	return names[rand.Intn(len(names))]
}

func getRandomDogAge() int {
	return rand.Intn(10) + 1
}

func getRandomDogBreed() string {
	breeds := []string{"Beagle", "Bulldog", "Chihuahua", "Dalmatian", "Poodle", "Retriever", "Terrier"}
	return breeds[rand.Intn(len(breeds))]
}
