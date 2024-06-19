package routes

import (
	"ramenGO/cmd/api/controllers"
	"ramenGO/internal/repositories"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, repository repositories.RamenRepository) {

	r.GET("/broths", func(ctx *gin.Context) {
		controllers.ListBroths(ctx, repository)
	})

	r.GET("/proteins", func(ctx *gin.Context) {
		controllers.ListProteins(ctx, repository)
	})

	r.POST("/orders", func(ctx *gin.Context) {
		controllers.PlaceOrder(ctx, repository)
	})

}
