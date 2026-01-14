package services

import (
	e "financialcontrol/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) Pay(ctx *gin.Context) (int, []e.ApiError) {
	transaction, status, errs := t.read(ctx)

	if len(errs) > 0 {
		return status, errs
	}

	errs = t.transactionsRepository.Pay(ctx, transaction.ID, !transaction.Paid)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}
