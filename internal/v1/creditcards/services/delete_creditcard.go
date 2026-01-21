package services

import (
	"financialcontrol/internal/constants"
	e "financialcontrol/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) Delete(ctx *gin.Context) (int, []e.ApiError) {
	creditcard, status, err := c.read(ctx)

	if len(err) > 0 {
		return status, err
	}

	hasTransactions, err := c.repository.HasTransactionsByCreditCard(ctx, creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	if hasTransactions {
		return http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.CreditcardCannotBeDeletedMsg}}
	}

	err = c.repository.DeleteCreditCard(ctx, creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
