package modelsdto

import (
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func CreateMonthlyTransactionFromRequest(model dtos.MonthlyTransactionRequest, userID uuid.UUID) models.CreateMonthlyTransaction {
	return models.CreateMonthlyTransaction{
		UserID:       userID,
		Name:         model.Name,
		Value:        model.Value,
		Day:          model.Day,
		CategoryID:   model.CategoryID,
		CreditCardID: model.CreditCardID,
	}
}

func MonthlyTransactionResponseFromModel(model models.MonthlyTransaction) dtos.MonthlyTransactionResponse {
	var creditcard *dtos.ShortCreditCardResponse
	if model.Creditcard != nil {
		creditcard = ShortCreditCardResponseFromShortCreditCard(model.Creditcard)
	}

	return dtos.MonthlyTransactionResponse{
		ID:         model.ID,
		Value:      model.Value,
		Day:        model.Day,
		Category:   ShortCategoryResponseFromShortModel(model.Category),
		Creditcard: creditcard,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}
