package use_cases

import (
	"ramenGO/internal/entities"
	"ramenGO/internal/repositories"
	errormessage "ramenGO/pkg/error-message"
)

type listBrothsUseCase struct {
	repository repositories.RamenRepository
}

func NewListBrothsUseCase(repo repositories.RamenRepository) *listBrothsUseCase {
	return &listBrothsUseCase{
		repository: repo,
	}
}

func (u *listBrothsUseCase) Execute() ([]*entities.Broth, *errormessage.ErrorResponse) {
	broths, err := u.repository.GetBroths()

	if err != nil {
		return nil, err
	}

	return broths, nil
}
