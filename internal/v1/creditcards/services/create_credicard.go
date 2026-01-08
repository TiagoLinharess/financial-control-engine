package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"
)

func (c CreditCardsService) Create(w http.ResponseWriter, r *http.Request) (cm.CreditCardResponse, int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.ReadCountByUser(r.Context(), userID)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return cm.CreditCardResponse{}, http.StatusForbidden, []e.ApiError{e.LimitError{Message: e.CreditcardsLimit}}
	}

	request, errs := u.DecodeValidJson[cm.CreditCardRequest](r)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	model := request.ToCreateModel(userID)

	creditCard, errs := c.repository.Create(r.Context(), model)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	return creditCard.ToResponse(), http.StatusCreated, nil
}
