package models

import (
	"financialcontrol/internal/models"
	"time"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID              uuid.UUID              `json:"id"`
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}
