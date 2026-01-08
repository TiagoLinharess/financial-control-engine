package services

import (
	e "financialcontrol/internal/models/errors"
	"net/http"
)

func (c CreditCardsService) Delete(w http.ResponseWriter, r *http.Request) (int, []e.ApiError) {
	creditcard, status, err := c.read(w, r)

	if len(err) > 0 {
		return status, err
	}

	err = c.repository.Delete(r.Context(), creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
