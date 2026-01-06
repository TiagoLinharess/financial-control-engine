package services

import (
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"
)

func (c CreditCardsService) Create(w http.ResponseWriter, r *http.Request) (cm.CreditCardResponse, int, []e.ApiError) {
	return cm.CreditCardResponse{}, http.StatusCreated, nil
}
