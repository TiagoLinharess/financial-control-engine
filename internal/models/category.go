package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID              uuid.UUID       `json:"id"`
	UserID          uuid.UUID       `json:"user_id"`
	TransactionType TransactionType `json:"transaction_type"`
	Name            string          `json:"name"`
	Icon            string          `json:"icon"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type CreateCategory struct {
	UserID          uuid.UUID
	TransactionType TransactionType
	Name            string
	Icon            string
}

type ShortCategory struct {
	ID              uuid.UUID
	TransactionType TransactionType
	Name            string
	Icon            string
}
