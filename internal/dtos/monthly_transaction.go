package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/utils"
	"time"
)

type MonthlyTransactionResponse struct {
	ID         string                   `json:"id"`
	Value      float64                  `json:"value"`
	Day        int64                    `json:"day"`
	Category   ShortCategoryResponse    `json:"category"`
	Creditcard *ShortCreditCardResponse `json:"creditcard"`
	CreatedAt  time.Time                `json:"created_at"`
	UpdatedAt  time.Time                `json:"updated_at"`
}

type MonthlyTransactionRequest struct {
	Name         string  `json:"name" binding:"required"`
	Value        float64 `json:"value" binding:"required"`
	Day          int     `json:"day" binding:"required"`
	CategoryID   string  `json:"category_id" binding:"required"`
	CreditCardID string  `json:"credit_card_id"`
}

func (m MonthlyTransactionRequest) Validate() []errors.ApiError {
	errs := make([]errors.ApiError, 0)

	if utils.IsBlank(m.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.MonthlyTransactionNameEmptyMsg})
	}

	if len(m.Name) < 2 || len(m.Name) > 255 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.MonthlyTransactionNameInvalidCharsCountMsg})
	}

	if m.Value < 0 || m.Value >= 1000000000000000.00 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.MonthlyTransactionValueInvalidMsg})
	}

	if m.Day < 1 || m.Day > 31 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.MonthlyTransactionDayInvalidMsg})
	}

	return errs
}
