package Configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB connection function
func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected successfully !!!! ")
	return client
}

// // MongoDB Client instance
var DB *mongo.Client = ConnectDB()

//Getting database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golang-masterclass").Collection(collectionName)
	return collection
}
