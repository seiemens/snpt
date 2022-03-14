/*
Created by Ramon
Date: 7.3.22
*/
package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "snpt/lib"
)

func GetSnippetByID(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, util.GetMongoSnippetByKey("id", id))
}

func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"lol": "lol"})
}

func CreateSnippet(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"lol": "soos"})
}

func CreateCookie(c *gin.Context) {
	cookie := c.Param("cookie")
	if cookie == "" {
		cookie = util.GenerateRandomString(69, false)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"cookie": cookie})
}
