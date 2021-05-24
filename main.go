package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/proph/geizhals-api-go/helpers"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Welcome to geizhals-api-go"})
	})

	server.Run(":" + helpers.GetEnv("PORT", "8080"))
}
