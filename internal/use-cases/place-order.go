package use_cases

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ramenGO/internal/entities"
	"ramenGO/internal/repositories"
	errormessage "ramenGO/pkg/error-message"
	"time"
)

type placeOrderUseCase struct {
	repository repositories.RamenRepository
}

func NewPlaceOrderUseCase(repo repositories.RamenRepository) *placeOrderUseCase {
	return &placeOrderUseCase{
		repository: repo,
	}
}

func (u *placeOrderUseCase) Execute(order *entities.Order) (*entities.OrderPlacedSuccess, *errormessage.ErrorResponse) {
	producList, err := u.repository.GetProductsList()

	if err != nil {
		return nil, err
	}

	if err := producList.CheckOrderIDs(order.BrothID, order.ProteinID); err != nil {
		return nil, err
	}

	broth := producList.BrothsIDs[order.BrothID].Name
	protein := producList.ProteinsIDs[order.ProteinID].Name

	orderSuccess, err := u.getPlaceOrder(broth, protein)

	if err != nil {
		return nil, err
	}

	return orderSuccess, nil
}
func (u *placeOrderUseCase) getPlaceOrder(broth, protein string) (*entities.OrderPlacedSuccess, *errormessage.ErrorResponse) {
	orderID, err := u.getOrderID()

	if err != nil {
		return nil, err
	}

	orderImageURL, _ := u.repository.GetOrderImage(protein)
	orderDescription := fmt.Sprintf("%s and %s Ramen", broth, protein)

	orderSuccess := &entities.OrderPlacedSuccess{
		ID:          *orderID,
		Description: orderDescription,
		Image:       orderImageURL}

	return orderSuccess, nil
}

func (u *placeOrderUseCase) getOrderID() (*string, *errormessage.ErrorResponse) {
	var orderID entities.OrderID
	errorResponse := &errormessage.ErrorResponse{Code: http.StatusInternalServerError, Message: "could not place order"}

	url := "https://api.tech.redventures.com.br/orders/generate-id"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, errorResponse
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-api-key", u.repository.GetAccessKey())

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errorResponse
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errorResponse
		}

		if err := json.Unmarshal(body, &orderID); err != nil {
			return nil, errorResponse
		}
	} else {
		return nil, errorResponse
	}

	return &orderID.ID, nil
}
