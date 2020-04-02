package api

import (
	"fmt"

	"github.com/stephencdaly/stephens-openbanking-test/database"
)

type PaymentResponse struct {
	PaymentID   string `json:"payment_id"`
	Reference   string `json:"reference"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	NextURL     string `json:"next_url"`
	Status      string `json:"status"`
}

func NewPaymentResponse(charge database.Charge) *PaymentResponse {
	response := PaymentResponse{
		PaymentID:   charge.ExternalId,
		Reference:   charge.Reference,
		Description: charge.Description,
		Amount:      charge.Amount,
		NextURL:     fmt.Sprintf("https://localhost:8080/payment/%s", charge.ExternalId),
		Status:      charge.Status,
	}
	return &response
}
