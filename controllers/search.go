package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/proph/geizhals-api-go/services"
)

func InitSearchController(router *gin.RouterGroup) {
	r := router.Group("/search")
	r.GET("/:query", getSearchResult(0))
}

func getSearchResult(category int) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		query := ctx.Param("query")
		result := services.Search(query, category)
		response := map[string]interface{}{
			"result": result,
		}
		ctx.JSON(http.StatusOK, response)
	}
}
