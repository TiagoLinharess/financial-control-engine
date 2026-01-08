package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) Update(ctx *gin.Context) (cm.CreditCardResponse, int, []e.ApiError) {
	creditcard, status, err := c.read(ctx)

	if len(err) > 0 {
		return cm.CreditCardResponse{}, status, err
	}

	request, errs := u.DecodeValidJson[cm.CreditCardRequest](ctx)

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

	creditcard, err = c.repository.Update(ctx, creditcard)

	return creditcard.ToResponse(), http.StatusOK, nil
}
