package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateTransaction struct {
	UserID                    uuid.UUID
	Name                      string
	Date                      time.Time
	Value                     float64
	CategoryID                uuid.UUID
	CreditcardID              *uuid.UUID
	MonthlyTransactionsID     *uuid.UUID
	AnnualTransactionsID      *uuid.UUID
	InstallmentTransactionsID *uuid.UUID
}
