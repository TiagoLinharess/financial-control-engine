package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) Read(ctx *gin.Context) ([]cm.CreditCardResponse, int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return []cm.CreditCardResponse{}, http.StatusUnauthorized, errs
	}

	creditCards, errs := c.repository.Read(ctx, userID)

	if len(errs) > 0 {
		return []cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	creditCardsResponse := make([]cm.CreditCardResponse, 0, len(creditCards))

	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, creditCard.ToResponse())
	}

	return creditCardsResponse, http.StatusOK, nil
}
