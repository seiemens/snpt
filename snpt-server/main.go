/*
Created by Ramon
Date: 7.3.22
*/
package main

import (
	"snpt/endpoints"
	"snpt/lib"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Connection URI

func main() {
	lib.ConnectToDb()

	router := gin.Default()
	router.GET("/s/:id", endpoints.GetSnippetByID)
	router.GET("/s", endpoints.GetSnippets)
	router.GET("/cookie", endpoints.CreateCookie)
	router.POST("/create", endpoints.CreateSnippet)
	router.POST("/edit", endpoints.EditSnippet)

	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	errGin := router.Run("localhost:3333")
	if errGin != nil {
		return
	}
}
