package models

import (
	"time"

	"github.com/google/uuid"
)

type CreditCardResponse struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
