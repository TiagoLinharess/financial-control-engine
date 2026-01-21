package models

import (
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

type ShortCategory struct {
	ID              uuid.UUID
	TransactionType models.TransactionType
	Name            string
	Icon            string
}

func (c ShortCategory) ToShortResponse() ShortCategoryResponse {
	return ShortCategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}
