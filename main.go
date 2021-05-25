package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/proph/geizhals-api-go/controllers"
	"github.com/proph/geizhals-api-go/helpers"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()

	defaultGroup := server.Group("/api")
	controllers.InitSearchController(defaultGroup)

	server.Run(":" + helpers.GetEnv("PORT", "8080"))
}
