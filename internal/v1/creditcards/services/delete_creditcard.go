package services

import (
	e "financialcontrol/internal/models/errors"
	"net/http"
)

func (c CreditCardsService) Delete(w http.ResponseWriter, r *http.Request) (int, []e.ApiError) {
	return http.StatusOK, nil
}
