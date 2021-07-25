package controllers

import (
	"net/http"

	"github.com/dschila/geizhals-api-go/services"
	"github.com/gin-gonic/gin"
)

// Init 'custom-filer' endpoints
func InitCustomFilterController(router *gin.RouterGroup) {
	r := router.Group("/custom-filter")
	r.GET("/:query", getResult())
}

// Returns the search result as JSON
func getResult() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query := ctx.Param("query")
		result := services.CustomFilter(query)
		ctx.JSON(http.StatusOK, result)
	}
}
