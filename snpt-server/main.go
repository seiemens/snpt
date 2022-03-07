package main

import (
	"snpt/endpoints"

	"github.com/gin-gonic/gin"
)

// Connection URI

func main() {
	router := gin.Default()
	router.GET("/s/:id", endpoints.GetSnippetByID)
	router.GET("/s", endpoints.GetSnippets)

	errGin := router.Run("localhost:3333")
	if errGin != nil {
		return
	}
}
