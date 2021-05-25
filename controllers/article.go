package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/proph/geizhals-api-go/services"
)

// Sets the endpoints below the Router-DefaultGroup
func InitArticleController(router *gin.RouterGroup) {
	r := router.Group("/article")
	r.GET("/:identifier", getArticleDetails())
}

// Return the article details including the all providers.
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
