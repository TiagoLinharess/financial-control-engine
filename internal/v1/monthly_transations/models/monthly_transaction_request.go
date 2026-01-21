package models

import (
	"financialcontrol/internal/constants"
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"

	"github.com/google/uuid"
)

type MonthlyTransactionRequest struct {
	Name         string  `json:"name" binding:"required"`
	Value        float64 `json:"value" binding:"required"`
	Day          int     `json:"day" binding:"required"`
	CategoryID   string  `json:"category_id" binding:"required"`
	CreditCardID string  `json:"credit_card_id"`
}

func (m MonthlyTransactionRequest) Validate() []e.ApiError {
	errs := make([]e.ApiError, 0)

	if u.IsBlank(m.Name) {
		errs = append(errs, e.InvalidFieldError{Message: constants.MonthlyTransactionNameEmptyMsg})
	}

	if len(m.Name) < 2 || len(m.Name) > 255 {
		errs = append(errs, e.InvalidFieldError{Message: constants.MonthlyTransactionNameInvalidCharsCountMsg})
	}

	if m.Value < 0 || m.Value >= 1000000000000000.00 {
		errs = append(errs, e.InvalidFieldError{Message: constants.MonthlyTransactionValueInvalidMsg})
	}

	if m.Day < 1 || m.Day > 31 {
		errs = append(errs, e.InvalidFieldError{Message: constants.MonthlyTransactionDayInvalidMsg})
	}

	return errs
}

func (m *MonthlyTransactionRequest) ToCreateModel(userID uuid.UUID) CreateMonthlyTransaction {
	return CreateMonthlyTransaction{
		UserID:       userID,
		Name:         m.Name,
		Value:        m.Value,
		Day:          m.Day,
		CategoryID:   m.CategoryID,
		CreditCardID: m.CreditCardID,
	}
}
