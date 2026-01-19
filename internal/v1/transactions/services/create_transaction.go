package services

import (
	e "financialcontrol/internal/models/errors"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) Create(ctx *gin.Context) (tm.TransactionResponse, int, []e.ApiError) {
	relations, status, errs := t.getRelations(ctx)

	if len(errs) > 0 {
		return tm.TransactionResponse{}, status, errs
	}

	data := relations.Request.ToCreateModel(relations.UserID)

	transaction, errs := t.transactionsRepository.Create(ctx, data)

	if len(errs) > 0 {
		return tm.TransactionResponse{}, http.StatusInternalServerError, errs
	}

	response := transaction.ToResponse(relations.CategoryResponse, relations.CreditcardResponse)

	return response, http.StatusCreated, nil
}
