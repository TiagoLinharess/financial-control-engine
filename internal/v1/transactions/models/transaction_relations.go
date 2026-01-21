package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

type TransactionRelations struct {
	UserID             uuid.UUID
	Request            TransactionRequest
	CategoryResponse   cm.ShortCategoryResponse
	CreditcardResponse *cr.ShortCreditCardResponse
}
