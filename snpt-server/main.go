/*
Created by Ramon
Date: 7.3.22
*/
package main

import (
	"snpt/endpoints"
	"snpt/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	lib.ConnectToDb()

	router := gin.New()
	router.Use(CORSMiddleware())

	router.GET("/s/:id", endpoints.GetSnippetByID)
	router.GET("/s", endpoints.GetSnippets)
	router.GET("/cookie", endpoints.CreateCookie)
	router.POST("/create", endpoints.CreateSnippet)
	router.POST("/edit", endpoints.EditSnippet)
	router.DELETE("/delete", endpoints.DeleteSnippet)

	errGin := router.Run("localhost:3333")
	if errGin != nil {
		return
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
