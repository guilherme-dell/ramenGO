package repositories

import (
	"ramenGO/internal/entities"
	errormessage "ramenGO/pkg/error-message"
)

type RamenRepository interface {
	GetBroths() ([]*entities.Broth, *errormessage.ErrorResponse)
	GetProteins() ([]*entities.Protein, *errormessage.ErrorResponse)
	GetProductsList() (*entities.ProductList, *errormessage.ErrorResponse)
	GetOrderImage(orderProtein string) (string, *errormessage.ErrorResponse)
	GetAccessKey() string
}
