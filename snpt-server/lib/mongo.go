/*
Created by Jordan
Date: 7.3.22
Functions:
	ConnectToDB
	CloseDBConnection
	GetMongoSnippetByKey
	GetAllSnippets
	DeleteAllSnippetsByCookie
	CreateSnippet
	EditSnippet
	DeleteSnippet
*/
package lib

//The Import obviously
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

//Global Variables with Data needed by functions below
var Uri = GoDotEnvVariable("MONGO_URL")
var Client *mongo.Client

//Establishes a Connection to the DB
//Values: none
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

//Closes the DB connection
//Values: The mongoDBclient connection valiable
func CloseDbConnection(client *mongo.Client) {
	var err error
	if err = client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

//Gets a Snippet with an encoded Key provided by Frontend
//Values: Key(encoded string), value(the value of the search)
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

//Gets all Snippets
//values: none
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

//Delete All snippets created by the User
//values: cookie(the User)
func DeleteAllSnippetsByCookie(cookie string) interface{} {
	for _, v := range GetAllSnippets() {
		if v.Cookie == cookie {
			DeleteSnippet(v.ID, v.Cookie)
		}
	}

	return "oge"
}

//Create a Snippet by the specific User
//Values: Title(Title of snippet), content(The content of the Snippet), Cookie(The User which has given the command to create the Snippet)
func CreateSnippet(title, content, cookie string) interface{} {

	var idExist bool = true
	var x models.Snippet
	for idExist {
		idExist = false
		x = models.Snippet{
			ID:      GenerateRandomString(6, false),
			Title:   title,
			Content: content,
			Cookie:  cookie,
		}
		snippets := GetAllSnippets()
		for _, value := range snippets {
			if value.ID == x.ID {
				idExist = true
			}
		}
	}

	snptDB := Client.Database("snpt")
	snippetCollection := snptDB.Collection("snippets")
	_, err := snippetCollection.InsertOne(context.Background(), x)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(x)
	return x
}

//Gets and Edits a snippet and saves it to DB again afterwards
//values: id(the id of the Snippet), title(the title of the Snippet), content(the body of the Snippet), cookie(the User)
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
		return "Updated Snippet"
	} else {
		return "You are most likely perhaps probably maybe possibly not the owner of this Item."
	}

}

//Deletes a snippet, that is in the DB, otherwise no action taken
//Values: id(the ID of the snippet), cookie(the User)
func DeleteSnippet(id, cookie string) string {
	snippetCollection := Client.Database("snpt").Collection("snippets")
	var shouldBeCookie string
	test := GetMongoSnippetByKey("id", id)
	for _, snippet := range test {
		shouldBeCookie = snippet.Cookie
	}

	if shouldBeCookie == cookie {

		_, err := snippetCollection.DeleteOne(context.Background(), bson.M{"id": id})
		if err != nil {
			log.Fatal(err)
		}
		return "Snippet has been deleted."
	}
	return "You are most likely perhaps probably maybe possibly not the owner of this Item."
}
