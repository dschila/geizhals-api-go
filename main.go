package main

import (
	"log"

	"github.com/dschila/geizhals-api-go/controllers"
	"github.com/dschila/geizhals-api-go/helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	server := gin.Default()
	server.Use(gin.Logger())
	server.Use(cors.Default())

	defaultGroup := server.Group("/api")
	controllers.InitSearchController(defaultGroup)
	controllers.InitArticleController(defaultGroup)
	controllers.InitCustomFilterController(defaultGroup)

	server.Run(":" + helpers.GetEnv("PORT", "8080"))
}
