package endpoints

import (
	"github.com/gin-gonic/gin"

	"net/http"
	
)

func GetSnippetByID(c *gin.Context) {
	//id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	//for _, a := range albums {
	//	if a.ID == id {
	//		c.IndentedJSON(http.StatusOK, a)
	//		return
	//	}
	//}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "snippet not found"})
}

func GetSnippets(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"lol": "lol"})
}
