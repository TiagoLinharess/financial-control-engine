package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionsCreditCardTotal struct {
	Date         time.Time
	UserID       uuid.UUID
	CreditcardID uuid.UUID
}
