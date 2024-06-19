package entities

import (
	"net/http"
	errormessage "ramenGO/pkg/error-message"
)

type Order struct {
	BrothID   string `json:"brothId" binding:"required"`
	ProteinID string `json:"proteinId" binding:"required"`
}

type OrderID struct {
	ID string `json:"orderId"`
}

type OrderPlacedSuccess struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Product struct {
	Status bool
	Name   string
}

type ProductList struct {
	BrothsIDs   map[string]Product
	ProteinsIDs map[string]Product
}

func (p *ProductList) CheckOrderIDs(brothID, pronteinID string) *errormessage.ErrorResponse {
	errorResponse := &errormessage.ErrorResponse{Code: http.StatusBadRequest, Message: "both brothId and proteinId are required"}

	if _, brothExists := p.BrothsIDs[brothID]; !brothExists {
		return errorResponse
	}

	if _, proteinsExists := p.ProteinsIDs[pronteinID]; !proteinsExists {
		return errorResponse
	}

	return nil
}
