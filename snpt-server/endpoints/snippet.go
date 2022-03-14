package endpoints

import (
	mongo "snpt/lib"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetSnippetByID(c *gin.Context) {
	id := c.Param("id")

	c.IndentedJSON(http.StatusOK, mongo.GetMongoSnippetByKey("id", id))
}

func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mongo.GetAllSnippets())
}

func CheckAndCreateCookie(c *gin.Context) {

}
