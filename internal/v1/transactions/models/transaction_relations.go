package models

import (
	"financialcontrol/internal/categories"
	cr "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

type TransactionRelations struct {
	UserID             uuid.UUID
	Request            TransactionRequest
	CategoryResponse   categories.ShortCategoryResponse
	CreditcardResponse *cr.ShortCreditCardResponse
}
