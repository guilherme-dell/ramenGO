package controllers

import (
	"net/http"
	"ramenGO/internal/repositories"
	use_cases "ramenGO/internal/use-cases"

	"github.com/gin-gonic/gin"
)

func ListProteins(ctx *gin.Context, repo repositories.RamenRepository) {
	useCase := use_cases.NewListProteinsUseCase(repo)

	proteins, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
	}

	ctx.JSON(http.StatusOK, proteins)
}
