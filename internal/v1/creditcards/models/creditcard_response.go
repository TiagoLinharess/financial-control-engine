package models

import (
	"time"

	"github.com/google/uuid"
)

type CreditCardResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
	BackgroundColor  string    `json:"background_color"`
	TextColor        string    `json:"text_color"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
