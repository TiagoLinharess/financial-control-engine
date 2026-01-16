package models

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionRequest struct {
	Name                      string     `json:"name"`
	Date                      time.Time  `json:"date"`
	Value                     float64    `json:"value"`
	Paid                      bool       `json:"paid"`
	CategoryID                uuid.UUID  `json:"category_id"`
	CreditcardID              *uuid.UUID `json:"creditcard_id,omitempty"`
	MonthlyTransactionsID     *uuid.UUID `json:"monthly_transactions_id,omitempty"`
	AnnualTransactionsID      *uuid.UUID `json:"annual_transactions_id,omitempty"`
	InstallmentTransactionsID *uuid.UUID `json:"installment_transactions_id,omitempty"`
}

func (t TransactionRequest) Validate() []errors.ApiError {
	errs := make([]errors.ApiError, 0)

	if utils.IsBlank(t.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionNameEmptyMsg})
	}

	if len(t.Name) > 255 || len(t.Name) < 2 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionNameInvalidCharsCountMsg})
	}

	if t.Value < 0 || t.Value >= 1000000000000000.00 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionAmountInvalidMsg})
	}

	// TODO: validate date
	// TODO: validate field sizes

	return errs
}

func (t TransactionRequest) ToCreateModel(userID uuid.UUID) CreateTransaction {
	return CreateTransaction{
		UserID:                    userID,
		Name:                      t.Name,
		Date:                      t.Date,
		Value:                     t.Value,
		Paid:                      t.Paid,
		CategoryID:                t.CategoryID,
		CreditcardID:              t.CreditcardID,
		MonthlyTransactionsID:     t.MonthlyTransactionsID,
		AnnualTransactionsID:      t.AnnualTransactionsID,
		InstallmentTransactionsID: t.InstallmentTransactionsID,
	}
}
