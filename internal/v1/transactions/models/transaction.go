package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                        uuid.UUID
	UserID                    uuid.UUID
	Name                      string
	Date                      time.Time
	Value                     float64
	CategoryID                uuid.UUID
	CreditCardID              *uuid.UUID
	MonthlyTransactionsID     *uuid.UUID
	AnnualTransactionsID      *uuid.UUID
	InstallmentTransactionsID *uuid.UUID
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func (t Transaction) ToResponse() TransactionResponse {
	return TransactionResponse{
		ID:                        t.ID,
		Name:                      t.Name,
		Date:                      t.Date,
		Value:                     t.Value,
		CategoryID:                t.CategoryID,
		CreditCardID:              t.CreditCardID,
		MonthlyTransactionsID:     t.MonthlyTransactionsID,
		AnnualTransactionsID:      t.AnnualTransactionsID,
		InstallmentTransactionsID: t.InstallmentTransactionsID,
		CreatedAt:                 t.CreatedAt,
		UpdatedAt:                 t.UpdatedAt,
	}
}
