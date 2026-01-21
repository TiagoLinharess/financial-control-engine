package models

import (
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

type ShortCategoryResponse struct {
	ID              uuid.UUID              `json:"id"`
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
}
