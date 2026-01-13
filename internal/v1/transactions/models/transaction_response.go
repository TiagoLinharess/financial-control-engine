package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	"time"

	"github.com/google/uuid"
)

type TransactionResponse struct {
	ID                     uuid.UUID                   `json:"id"`
	Name                   string                      `json:"name"`
	Date                   time.Time                   `json:"date"`
	Value                  float64                     `json:"value"`
	Category               cm.ShortCategoryReponse     `json:"category"`
	Creditcard             *cr.ShortCreditCardResponse `json:"creditcard"`
	MonthlyTransaction     *uuid.UUID                  `json:"monthly_transaction,omitempty"`
	AnnualTransaction      *uuid.UUID                  `json:"annual_transaction,omitempty"`
	InstallmentTransaction *uuid.UUID                  `json:"installment_transaction,omitempty"`
	CreatedAt              time.Time                   `json:"created_at"`
	UpdatedAt              time.Time                   `json:"updated_at"`
}
