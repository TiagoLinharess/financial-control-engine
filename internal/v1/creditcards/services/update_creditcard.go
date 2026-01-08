package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"
)

func (c CreditCardsService) Update(w http.ResponseWriter, r *http.Request) (cm.CreditCardResponse, int, []e.ApiError) {
	creditcard, status, err := c.read(w, r)

	if len(err) > 0 {
		return cm.CreditCardResponse{}, status, err
	}

	request, errs := u.DecodeValidJson[cm.CreditCardRequest](r)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	creditcard.Name = request.Name
	creditcard.FirstFourNumbers = request.FirstFourNumbers
	creditcard.Limit = request.Limit
	creditcard.CloseDay = request.CloseDay
	creditcard.ExpireDay = request.ExpireDay
	creditcard.BackgroundColor = request.BackgroundColor
	creditcard.TextColor = request.TextColor

	creditcard, err = c.repository.Update(r.Context(), creditcard)

	return creditcard.ToResponse(), http.StatusOK, nil
}
