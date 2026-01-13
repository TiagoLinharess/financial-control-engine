package models

import "github.com/google/uuid"

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

func (c *ShortCreditCard) ToShortResponse() *ShortCreditCardResponse {
	return &ShortCreditCardResponse{
		ID:               c.ID,
		Name:             c.Name,
		FirstFourNumbers: c.FirstFourNumbers,
		Limit:            c.Limit,
		CloseDay:         c.CloseDay,
		ExpireDay:        c.ExpireDay,
		BackgroundColor:  c.BackgroundColor,
		TextColor:        c.TextColor,
	}
}
