package services

import (
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/categories/models"
	crm "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) Update(ctx *gin.Context) (tm.TransactionResponse, int, []e.ApiError) {
	relations, status, errs := t.getRelations(ctx)

	if len(errs) > 0 {
		return tm.TransactionResponse{}, status, errs
	}

	transaction, status, errs := t.read(ctx)

	if len(errs) > 0 {
		return tm.TransactionResponse{}, status, errs
	}

	var creditcard *crm.ShortCreditCard
	if relations.CreditcardResponse != nil {
		creditcardModel := crm.ShortCreditCard(*relations.CreditcardResponse)
		creditcard = &creditcardModel
	}

	transaction.Name = relations.Request.Name
	transaction.Date = relations.Request.Date
	transaction.Paid = relations.Request.Paid
	transaction.Category = cm.ShortCategory(relations.CategoryResponse)
	transaction.Creditcard = creditcard

	transactionUpdated, errs := t.transactionsRepository.Update(ctx, transaction)

	if len(errs) > 0 {
		return tm.TransactionResponse{}, http.StatusInternalServerError, errs
	}

	response := transactionUpdated.ToResponse(relations.CategoryResponse, relations.CreditcardResponse)

	return response, http.StatusOK, nil
}
