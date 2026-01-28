package models

import (
	"time"

	"github.com/google/uuid"
)

type MonthlyTransaction struct {
	ID         string
	UserID     uuid.UUID
	Value      float64
	Day        int64
	Category   ShortCategory
	Creditcard *ShortCreditCard
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateMonthlyTransaction struct {
	UserID       uuid.UUID
	Name         string
	Value        float64
	Day          int
	CategoryID   string
	CreditCardID string
}
