package services

import (
	e "financialcontrol/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CreditCardsService) Delete(ctx *gin.Context) (int, []e.ApiError) {
	creditcard, status, err := c.read(ctx)

	if len(err) > 0 {
		return status, err
	}

	// TODO: check if creditcard has transactions associated

	err = c.repository.Delete(ctx, creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
