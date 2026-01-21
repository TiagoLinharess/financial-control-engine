package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	"time"
)

type MonthlyTransactionResponse struct {
	ID         string              `json:"id"`
	Value      float64             `json:"value"`
	Day        int64               `json:"day"`
	Category   cm.ShortCategory    `json:"category"`
	Creditcard *cr.ShortCreditCard `json:"creditcard"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}
