package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionRequest struct {
	Name                      string     `json:"name"`
	Date                      time.Time  `json:"date"`
	Value                     float64    `json:"value"`
	CategoryID                uuid.UUID  `json:"category_id"`
	CreditCardID              *uuid.UUID `json:"credit_card_id,omitempty"`
	MonthlyTransactionsID     *uuid.UUID `json:"monthly_transactions_id,omitempty"`
	AnnualTransactionsID      *uuid.UUID `json:"annual_transactions_id,omitempty"`
	InstallmentTransactionsID *uuid.UUID `json:"installment_transactions_id,omitempty"`
}

func (t TransactionRequest) ToCreateModel(userID uuid.UUID) CreateTransaction {
	return CreateTransaction{
		UserID:                    userID,
		Name:                      t.Name,
		Date:                      t.Date,
		Value:                     t.Value,
		CategoryID:                t.CategoryID,
		CreditCardID:              t.CreditCardID,
		MonthlyTransactionsID:     t.MonthlyTransactionsID,
		AnnualTransactionsID:      t.AnnualTransactionsID,
		InstallmentTransactionsID: t.InstallmentTransactionsID,
	}
}
