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

var Uri = GoDotEnvVariable("MONGO_URL")
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

func GetAllSnippets() []models.Snippet {
	snptCollection := *Client.Database("snpt").Collection("snippets")
	cursor, err := snptCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var snippets []models.Snippet
	if err = cursor.All(context.Background(), &snippets); err != nil {
		log.Fatal(err)
	}
	return snippets
}

func CreateSnippet(title, content, cookie string) interface{} {
	x := models.Snippet{
		ID:      GenerateRandomString(6, false),
		Title:   title,
		Content: content,
		Cookie:  cookie,
	}

	snptDB := Client.Database("snpt")
	snippetCollection := snptDB.Collection("snippets")
	res, _ := snippetCollection.InsertOne(context.Background(), x)
	return res.InsertedID
}

func EditSnippet(id, title, content, cookie string) string {

	snippetCollection := Client.Database("snpt").Collection("snippets")
	var shouldBeCookie string
	test := GetMongoSnippetByKey("id", id)
	for _, snippet := range test {
		shouldBeCookie = snippet.Cookie
	}

	if shouldBeCookie == cookie {
		filter := bson.M{"id": bson.M{"$eq": id}}
		update := bson.M{"$set": bson.M{"title": title, "content": content}}
		_, err := snippetCollection.UpdateOne(
			context.Background(),
			filter,
			update,
		)
		if err != nil {
			log.Fatal(err)
		}
		return "Updated Document"
	} else {
		return "You are most likely perhaps probably maybe possibly not the owner of this Item."
	}

}
