package services

import (
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/creditcards/models"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) ReadAt(ctx *gin.Context) (cm.CreditCardResponse, int, []e.ApiError) {
	creditcard, status, err := c.read(ctx)
	return creditcard.ToResponse(), status, err
}
