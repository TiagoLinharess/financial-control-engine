package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateCreditCard struct {
	UserID           uuid.UUID
	Name             string
	FirstFourNumbers string
	Limit            float64
	CloseDay         int32
	ExpireDay        int32
	BackgroundColor  string
	TextColor        string
}

type CreditCard struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	Name             string
	FirstFourNumbers string
	Limit            float64
	CloseDay         int32
	ExpireDay        int32
	BackgroundColor  string
	TextColor        string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ShortCreditCard struct {
	ID               uuid.UUID
	Name             string
	FirstFourNumbers string
	Limit            float64
	CloseDay         int32
	ExpireDay        int32
	BackgroundColor  string
	TextColor        string
}
