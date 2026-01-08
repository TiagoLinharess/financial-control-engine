package models

import (
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
