package models

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"

	"github.com/google/uuid"
)

type CreditCardRequest struct {
	Name             string  `json:"name"`
	FirstFourNumbers string  `json:"first_four_numbers"`
	Limit            float64 `json:"limit"`
	CloseDay         int32   `json:"close_day"`
	ExpireDay        int32   `json:"expire_day"`
	BackgroundColor  string  `json:"background_color"`
	TextColor        string  `json:"text_color"`
}

func (c CreditCardRequest) Validate() []errors.ApiError {
	errs := make([]errors.ApiError, 0)

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: "Name must not be empty"})
	}

	if len(c.FirstFourNumbers) != 4 {
		errs = append(errs, errors.InvalidFieldError{Message: "First four numbers must have 4 characters"})
	}

	if c.Limit <= 0 {
		errs = append(errs, errors.InvalidFieldError{Message: "Limit must be greater than zero"})
	}

	if c.CloseDay < 1 || c.CloseDay > 31 {
		errs = append(errs, errors.InvalidFieldError{Message: "Close day must be between 1 and 31"})
	}

	if c.ExpireDay < 1 || c.ExpireDay > 31 {
		errs = append(errs, errors.InvalidFieldError{Message: "Expire day must be between 1 and 31"})
	}

	if utils.IsBlank(c.BackgroundColor) {
		errs = append(errs, errors.InvalidFieldError{Message: "Background color must not be empty"})
	}

	if utils.IsBlank(c.TextColor) {
		errs = append(errs, errors.InvalidFieldError{Message: "Text color must not be empty"})
	}

	return errs
}

func (c CreditCardRequest) ToCreateModel(userID uuid.UUID) CreateCreditCard {
	return CreateCreditCard{
		UserID:           userID,
		Name:             c.Name,
		FirstFourNumbers: c.FirstFourNumbers,
		Limit:            c.Limit,
		CloseDay:         c.CloseDay,
		ExpireDay:        c.ExpireDay,
		BackgroundColor:  c.BackgroundColor,
		TextColor:        c.TextColor,
	}
}
