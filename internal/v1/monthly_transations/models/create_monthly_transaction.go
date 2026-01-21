package models

import "github.com/google/uuid"

type CreateMonthlyTransaction struct {
	UserID       uuid.UUID
	Name         string
	Value        float64
	Day          int
	CategoryID   string
	CreditCardID string
}
