/*
Created by Ramon
Date: 7.3.22
*/
package endpoints

import (
	"fmt"
	mongo "snpt/lib"
	"snpt/models"

	"github.com/gin-gonic/gin"

	"net/http"
	util "snpt/lib"
)

func GetSnippetByID(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, util.GetMongoSnippetByKey("id", id))
}

func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mongo.GetAllSnippets())
}

func CreateSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.CreateSnippet(snpt.Title, snpt.Content, snpt.Cookie)})
}

func CreateCookie(c *gin.Context) {
	cookie := c.Param("cookie")
	if cookie == "" {
		cookie = util.GenerateRandomString(69, false)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
}
func EditSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.EditSnippet(snpt.ID, snpt.Title, snpt.Content, snpt.Cookie)})
}

func DeleteSnippet(c *gin.Context) {
	var snpt models.Snippet
	err := c.BindJSON(&snpt)
	if err != nil {
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.DeleteSnippet(snpt.ID, snpt.Cookie)})
}
