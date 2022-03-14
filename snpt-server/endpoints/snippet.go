/*
Created by Ramon
Date: 7.3.22
*/
package endpoints

import (
	mongo "snpt/lib"

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
	cookie := c.Param("cookie")
	title := c.Param("title")
	content := c.Param("content")
	c.IndentedJSON(http.StatusOK, gin.H{"answer": util.CreateSnippet(title, content, cookie)})
}

func CreateCookie(c *gin.Context) {
	cookie := c.Param("cookie")
	if cookie == "" {
		cookie = util.GenerateRandomString(69, false)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
}
