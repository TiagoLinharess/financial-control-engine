package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) Create(ctx *gin.Context) (cm.CreditCardResponse, int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.ReadCountByUser(ctx, userID)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return cm.CreditCardResponse{}, http.StatusForbidden, []e.ApiError{e.LimitError{Message: e.CreditcardsLimit}}
	}

	request, errs := u.DecodeValidJson[cm.CreditCardRequest](ctx)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	model := request.ToCreateModel(userID)

	creditCard, errs := c.repository.CreateCreditCard(ctx, model)

	if len(errs) > 0 {
		return cm.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	return creditCard.ToResponse(), http.StatusCreated, nil
}
