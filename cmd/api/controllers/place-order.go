package controllers

import (
	"net/http"
	"ramenGO/internal/entities"
	"ramenGO/internal/repositories"
	use_cases "ramenGO/internal/use-cases"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(ctx *gin.Context, repo repositories.RamenRepository) {
	var order entities.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "both brothId and proteinId are required"})
		return
	}

	useCase := use_cases.NewPlaceOrderUseCase(repo)

	orderSuccess, err := useCase.Execute(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusCreated, orderSuccess)
}
