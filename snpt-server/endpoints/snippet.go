package endpoints

import (
	"github.com/gin-gonic/gin"
	mongo "snpt/lib"

	"net/http"
)

func GetSnippetByID(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, mongo.GetMongoSnippetByKey("id", id))
}

func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"lol": "lol"})
}

func CheckAndCreateCookie(c *gin.Context) {

}
