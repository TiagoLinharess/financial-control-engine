package services

import (
	e "financialcontrol/internal/models/errors"
	s "financialcontrol/internal/store"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c CreditCardsService) ReadAt(w http.ResponseWriter, r *http.Request) (cm.CreditCardResponse, int, []e.ApiError) {
	creditcardNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CreditcardNotFound}}
	userID, errs := u.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusUnauthorized, errs
	}

	creditcardIdString := chi.URLParam(r, "id")

	creditcardId, err := uuid.Parse(creditcardIdString)

	if err != nil {
		return cm.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	creditCard, errs := c.repository.ReadByID(r.Context(), creditcardId)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == string(s.ErrNoRows)
		})
		if isNotFoundErr {
			return cm.CreditCardResponse{}, http.StatusNotFound, creditcardNotFoundErr
		}
		return cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	if creditCard.UserID != userID {
		return cm.CreditCardResponse{}, http.StatusForbidden, creditcardNotFoundErr
	}

	return creditCard.ToResponse(), http.StatusOK, nil
}
