package controllers

import (
	"net/http"
	"strconv"

	"github.com/dschila/geizhals-api-go/helpers"
	"github.com/dschila/geizhals-api-go/services"
	"github.com/gin-gonic/gin"
)

// Init 'search' endpoints
func InitSearchController(router *gin.RouterGroup) {
	r := router.Group("/search")
	r.GET("/:query", getSearchResult())
}

// Returns the search result as JSON
func getSearchResult() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query := ctx.Param("query")
		category, err := strconv.Atoi(ctx.DefaultQuery("category", "0"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(http.StatusBadRequest, err))
		} else {
			result := services.Search(query, category)
			ctx.JSON(http.StatusOK, result)
		}
	}
}
