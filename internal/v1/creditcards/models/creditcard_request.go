package models

import "github.com/google/uuid"

type CreditCardRequest struct {
	UserID           uuid.UUID `json:"user_id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
}
