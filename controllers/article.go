package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dschila/geizhals-api-go/services"
)

// Init 'article' endpoints
func InitArticleController(router *gin.RouterGroup) {
	r := router.Group("/article")
	r.GET("/:identifier", getArticleDetails())
}

// Returns the item details as JSON
func getArticleDetails() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		identifier := ctx.Param("identifier")
		article, provider := services.ArticleDetails(identifier)
		response := map[string]interface{}{
			"article":  article,
			"provider": provider,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
