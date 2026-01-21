package models

import (
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	"time"
)

type MonthlyTransaction struct {
	ID         string
	UserID     string
	Value      float64
	Day        int64
	Category   cm.ShortCategory
	Creditcard *cr.ShortCreditCard
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
