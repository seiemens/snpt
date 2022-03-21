/*
Created by Ramon
Date: 7.3.22
Functions:
	GetSnippedByID
	GetSnippets
	CreateSnippet
	CreateCookie
	EditSnippet
	DeleteSnippet
	DeleteAllSnippets
*/
package endpoints

//Package imports obviously
import (
	"fmt"
	mongo "snpt/lib"
	"snpt/models"

	"github.com/gin-gonic/gin"

	"net/http"
	util "snpt/lib"
)

//Gets a snippet from the DB with a provided ID. Uses Gin
func GetSnippetByID(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, util.GetMongoSnippetByKey("id", id))
}

//Gets all snippets from the DB
func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mongo.GetAllSnippets())
}

//Creates a snippet with data provided by the Frontend
func CreateSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.CreateSnippet(snpt.Title, snpt.Content, snpt.Cookie)})
}

//creates a Cookie for each User, so each User can only Access his own snippets
func CreateCookie(c *gin.Context) {
	cookie := c.Param("cookie")
	if cookie == "" {
		cookie = util.GenerateRandomString(69, false)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
}

//Provides the functionality for Editing a Snippet with an Id provided by the Frontend
func EditSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.EditSnippet(snpt.ID, snpt.Title, snpt.Content, snpt.Cookie)})
}

//Deletes a Snippet whose ID is provided by the Frontend
func DeleteSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.DeleteSnippet(snpt.ID, snpt.Cookie)})
}

//Deletes all Snippets, that are in the DB
func DeleteAllSnippets(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.DeleteAllSnippetsByCookie(snpt.Cookie)})
}
