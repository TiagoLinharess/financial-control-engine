package models

import (
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

type CreateCategory struct {
	UserID          uuid.UUID
	TransactionType models.TransactionType
	Name            string
	Icon            string
}
