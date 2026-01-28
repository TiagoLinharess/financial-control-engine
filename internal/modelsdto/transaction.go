package modelsdto

import (
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func TransactionResponseFromShortTransaction(model models.ShortTransaction, category dtos.ShortCategoryResponse, creditcard *dtos.ShortCreditCardResponse) dtos.TransactionResponse {
	return dtos.TransactionResponse{
		ID:                     model.ID,
		Name:                   model.Name,
		Date:                   model.Date,
		Value:                  model.Value,
		Paid:                   model.Paid,
		Category:               category,
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              model.CreatedAt,
		UpdatedAt:              model.UpdatedAt,
	}
}

func CreateTransactionFromTransactionRequest(model dtos.TransactionRequest, userID uuid.UUID) models.CreateTransaction {
	return models.CreateTransaction{
		UserID:                    userID,
		Name:                      model.Name,
		Date:                      model.Date,
		Value:                     model.Value,
		Paid:                      model.Paid,
		CategoryID:                model.CategoryID,
		CreditcardID:              model.CreditcardID,
		MonthlyTransactionsID:     model.MonthlyTransactionsID,
		AnnualTransactionsID:      model.AnnualTransactionsID,
		InstallmentTransactionsID: model.InstallmentTransactionsID,
	}
}

func TransactionResponseFromTransaction(model models.Transaction) dtos.TransactionResponse {
	var creditcard *dtos.ShortCreditCardResponse
	if model.Creditcard != nil {
		creditcard = ShortCreditCardResponseFromShortCreditCard(model.Creditcard)
	}

	return dtos.TransactionResponse{
		ID:                     model.ID,
		Name:                   model.Name,
		Date:                   model.Date,
		Value:                  model.Value,
		Paid:                   model.Paid,
		Category:               ShortCategoryResponseFromShortModel(model.Category),
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              model.CreatedAt,
		UpdatedAt:              model.UpdatedAt,
	}
}
