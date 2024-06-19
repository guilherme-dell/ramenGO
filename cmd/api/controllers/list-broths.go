package controllers

import (
	"net/http"
	"ramenGO/internal/repositories"
	use_cases "ramenGO/internal/use-cases"

	"github.com/gin-gonic/gin"
)

func ListBroths(ctx *gin.Context, repo repositories.RamenRepository) {
	useCase := use_cases.NewListBrothsUseCase(repo)

	broths, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, broths)
}
