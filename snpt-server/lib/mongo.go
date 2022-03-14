/*
Created by Jordan
Date: 7.3.22
*/
package lib

import (
	"context"
	"fmt"
	"log"
	"snpt/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Uri string = goDotEnvVariable("MONGO_URL")
var Client *mongo.Client

func ConnectToDb() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	Client = client
	fmt.Println("Successfully connected and pinged.")
}

func CloseDbConnection(client *mongo.Client) {
	var err error
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func GetMongoSnippetByKey(key string, value interface{}) []models.Snippet {
	snptCollection := *Client.Database("snpt").Collection("snippets")
	res, err := snptCollection.Find(context.Background(), bson.D{{key, primitive.Regex{Pattern: "^.*" + value.(string) + ".*", Options: ""}}})
	if err != nil {
		log.Fatal(err)
	}
	var snippets []models.Snippet
	if err = res.All(context.Background(), &snippets); err != nil {
		log.Fatal(err)
	}
	return snippets
}
