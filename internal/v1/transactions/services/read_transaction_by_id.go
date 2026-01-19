package services

import (
	e "financialcontrol/internal/models/errors"
	tm "financialcontrol/internal/v1/transactions/models"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) ReadById(ctx *gin.Context) (tm.TransactionResponse, int, []e.ApiError) {
	transaction, status, errs := t.read(ctx)
	return transaction.ToResponse(), status, errs
}
