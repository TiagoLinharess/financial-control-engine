package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	"time"

	"github.com/google/uuid"
)

type ShortTransaction struct {
	ID        uuid.UUID
	Name      string
	Date      time.Time
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t ShortTransaction) ToResponse(category cm.ShortCategoryReponse, creditcard *cr.ShortCreditCardResponse) TransactionResponse {
	return TransactionResponse{
		ID:                     t.ID,
		Name:                   t.Name,
		Date:                   t.Date,
		Value:                  t.Value,
		Category:               category,
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              t.CreatedAt,
		UpdatedAt:              t.UpdatedAt,
	}
}
