package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type TransactionRelations struct {
	UserID             uuid.UUID
	Request            TransactionRequest
	CategoryResponse   ShortCategoryResponse
	CreditcardResponse *ShortCreditCardResponse
}

type TransactionResponse struct {
	ID                     uuid.UUID                `json:"id"`
	Name                   string                   `json:"name"`
	Date                   time.Time                `json:"date"`
	Value                  float64                  `json:"value"`
	Paid                   bool                     `json:"paid"`
	Category               ShortCategoryResponse    `json:"category"`
	Creditcard             *ShortCreditCardResponse `json:"creditcard"`
	MonthlyTransaction     *uuid.UUID               `json:"monthly_transaction,omitempty"`
	AnnualTransaction      *uuid.UUID               `json:"annual_transaction,omitempty"`
	InstallmentTransaction *uuid.UUID               `json:"installment_transaction,omitempty"`
	CreatedAt              time.Time                `json:"created_at"`
	UpdatedAt              time.Time                `json:"updated_at"`
}

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

	if len(t.Name) >= 255 || len(t.Name) <= 2 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionNameInvalidCharsCountMsg})
	}

	if t.Value < 0 || t.Value >= 1000000000000000.00 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionAmountInvalidMsg})
	}

	if t.Date.IsZero() {
		errs = append(errs, errors.InvalidFieldError{Message: constants.TransactionDateEmptyMsg})
	}

	return errs
}
