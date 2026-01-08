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

func (c CreateCreditCard) ToModel(id uuid.UUID, createdAt time.Time, updatedAt time.Time) CreditCard {
	return CreditCard{
		ID:               id,
		UserID:           c.UserID,
		Name:             c.Name,
		FirstFourNumbers: c.FirstFourNumbers,
		Limit:            c.Limit,
		CloseDay:         c.CloseDay,
		ExpireDay:        c.ExpireDay,
		BackgroundColor:  c.BackgroundColor,
		TextColor:        c.TextColor,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}
}
