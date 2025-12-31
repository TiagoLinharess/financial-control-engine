package models

import "financialcontrol/internal/models"

type CategoryRequest struct {
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
}
