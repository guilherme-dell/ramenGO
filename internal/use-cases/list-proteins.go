package use_cases

import (
	"ramenGO/internal/entities"
	"ramenGO/internal/repositories"
	errormessage "ramenGO/pkg/error-message"
)

type listProteinsUseCase struct {
	repository repositories.RamenRepository
}

func NewListProteinsUseCase(repo repositories.RamenRepository) *listProteinsUseCase {
	return &listProteinsUseCase{
		repository: repo,
	}
}

func (u *listProteinsUseCase) Execute() ([]*entities.Protein, *errormessage.ErrorResponse) {
	proteins, err := u.repository.GetProteins()

	if err != nil {
		return nil, err
	}

	return proteins, nil
}
