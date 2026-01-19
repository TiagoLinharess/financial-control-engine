package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                     uuid.UUID
	UserID                 uuid.UUID
	Name                   string
	Date                   time.Time
	Value                  float64
	Paid                   bool
	Category               cm.ShortCategory
	Creditcard             *cr.ShortCreditCard
	MonthlyTransaction     *uuid.UUID
	AnnualTransaction      *uuid.UUID
	InstallmentTransaction *uuid.UUID
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func (t Transaction) ToResponse() TransactionResponse {
	var creditcard *cr.ShortCreditCardResponse
	if t.Creditcard != nil {
		creditcard = t.Creditcard.ToShortResponse()
	}

	return TransactionResponse{
		ID:                     t.ID,
		Name:                   t.Name,
		Date:                   t.Date,
		Value:                  t.Value,
		Paid:                   t.Paid,
		Category:               t.Category.ToShortResponse(),
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              t.CreatedAt,
		UpdatedAt:              t.UpdatedAt,
	}
}
