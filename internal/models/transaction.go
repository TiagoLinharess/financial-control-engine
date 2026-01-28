package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID                     uuid.UUID
	UserID                 uuid.UUID
	Name                   string
	Date                   time.Time
	Value                  float64
	Paid                   bool
	Category               ShortCategory
	Creditcard             *ShortCreditCard
	MonthlyTransaction     *uuid.UUID
	AnnualTransaction      *uuid.UUID
	InstallmentTransaction *uuid.UUID
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type CreateTransaction struct {
	UserID                    uuid.UUID
	Name                      string
	Date                      time.Time
	Value                     float64
	Paid                      bool
	CategoryID                uuid.UUID
	CreditcardID              *uuid.UUID
	MonthlyTransactionsID     *uuid.UUID
	AnnualTransactionsID      *uuid.UUID
	InstallmentTransactionsID *uuid.UUID
}

type ShortTransaction struct {
	ID        uuid.UUID
	Name      string
	Date      time.Time
	Value     float64
	Paid      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionsCreditCardTotal struct {
	Date         time.Time
	UserID       uuid.UUID
	CreditcardID uuid.UUID
}
