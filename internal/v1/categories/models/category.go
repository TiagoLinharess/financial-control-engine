package models

import (
	"financialcontrol/internal/models"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID              uuid.UUID              `json:"id"`
	UserID          uuid.UUID              `json:"user_id"`
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

func (c Category) ToResponse() CategoryResponse {
	return CategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

func (c Category) ToShortResponse() ShortCategoryResponse {
	return ShortCategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}
