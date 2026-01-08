package models

import (
	"time"

	"github.com/google/uuid"
)

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

func (c CreditCard) ToResponse() CreditCardResponse {
	return CreditCardResponse{
		ID:               c.ID,
		UserID:           c.UserID,
		Name:             c.Name,
		FirstFourNumbers: c.FirstFourNumbers,
		Limit:            c.Limit,
		CloseDay:         c.CloseDay,
		ExpireDay:        c.ExpireDay,
		BackgroundColor:  c.BackgroundColor,
		TextColor:        c.TextColor,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
	}
}
