package services

import (
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"
)

func (c CreditCardsService) ReadAt(w http.ResponseWriter, r *http.Request) (cm.CreditCardResponse, int, []e.ApiError) {
	creditcard, status, err := c.read(w, r)
	return creditcard.ToResponse(), status, err
}
