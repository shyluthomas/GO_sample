package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func defaultPort(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "APP running on port: 8080")
}

func authMiddlware() gin.HandlerFunc {
	log.Printf("inside the middlwr")
	return func(c *gin.Context) {
		log.Printf("inside the middlwr")
		c.Next()
	}
}

func setupRoutes() {
	router := gin.Default()
	router.Use(authMiddlware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", defaultPort)
	router.GET("/albums", getAlbums)

	router.POST("/albums", postAlbum)

	router.GET("/album/:id", getAlbum)
	router.Run("localhost:8080")
}
