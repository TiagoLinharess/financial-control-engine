package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionResponse struct {
	ID                        uuid.UUID  `json:"id"`
	Name                      string     `json:"name"`
	Date                      time.Time  `json:"date"`
	Value                     float64    `json:"value"`
	CategoryID                uuid.UUID  `json:"category_id"`
	CreditCardID              *uuid.UUID `json:"credit_card_id"`
	MonthlyTransactionsID     *uuid.UUID `json:"monthly_transactions_id,omitempty"`
	AnnualTransactionsID      *uuid.UUID `json:"annual_transactions_id,omitempty"`
	InstallmentTransactionsID *uuid.UUID `json:"installment_transactions_id,omitempty"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}
