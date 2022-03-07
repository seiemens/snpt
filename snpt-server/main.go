package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"snpt/endpoints"
	"snpt/lib"
)

// Connection URI

func main() {
	lib.ConnectToDb()
	snippets := lib.GetMongoSnippetByKey("username", "eter")
	for _, s := range snippets {
		fmt.Println(s.Cookie)
	}

	//insert sample data

	//x := models.Snippet{
	//	ID:      lib.GenerateRandomString(6, false),
	//	Title:   "Mongo Du Mongo",
	//	Content: "Wieso gids dich eigentlich. wieso isch dis find so \"nice\" am been?",
	//	Cookie:  "pretendtoberandom",
	//}

	//snptDB := lib.Client.Database("snpt")
	//snippetCollection := snptDB.Collection("snippets")
	//snippetResult, err := snippetCollection.InsertOne(context.TODO(), x)
	//fmt.Println(snippetResult, err)

	router := gin.Default()
	router.GET("/s/:id", endpoints.GetSnippetByID)
	router.GET("/s", endpoints.GetSnippets)
	router.GET("/cookie", endpoints.CheckAndCreateCookie)

	errGin := router.Run("localhost:3333")
	if errGin != nil {
		return
	}
}
