package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ShortCreditCardResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
	BackgroundColor  string    `json:"background_color"`
	TextColor        string    `json:"text_color"`
}

type CreditCardResponse struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	FirstFourNumbers string    `json:"first_four_numbers"`
	Limit            float64   `json:"limit"`
	CloseDay         int32     `json:"close_day"`
	ExpireDay        int32     `json:"expire_day"`
	BackgroundColor  string    `json:"background_color"`
	TextColor        string    `json:"text_color"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreditCardRequest struct {
	Name             string  `json:"name"`
	FirstFourNumbers string  `json:"first_four_numbers"`
	Limit            float64 `json:"limit"`
	CloseDay         int32   `json:"close_day"`
	ExpireDay        int32   `json:"expire_day"`
	BackgroundColor  string  `json:"background_color"`
	TextColor        string  `json:"text_color"`
}

func (c CreditCardRequest) Validate() []errors.ApiErrorItem {
	errs := make([]errors.ApiErrorItem, 0)

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError(constants.NameEmptyMsg))
	}

	if len(c.Name) >= 100 || len(c.Name) <= 2 {
		errs = append(errs, errors.InvalidFieldError(constants.NameInvalidCharsCountMsg))
	}

	if len(c.FirstFourNumbers) != 4 {
		errs = append(errs, errors.InvalidFieldError(constants.FirstFourNumbersInvalidMsg))
	}

	if c.Limit <= 0 {
		errs = append(errs, errors.InvalidFieldError(constants.LimitInvalidMsg))
	}

	if c.CloseDay < 1 || c.CloseDay > 31 {
		errs = append(errs, errors.InvalidFieldError(constants.ClosingDayInvalidMsg))
	}

	if c.ExpireDay < 1 || c.ExpireDay > 31 {
		errs = append(errs, errors.InvalidFieldError(constants.ExpireDayInvalidMsg))
	}

	if utils.IsBlank(c.BackgroundColor) {
		errs = append(errs, errors.InvalidFieldError(constants.BackgroundColorEmptyMsg))
	}

	if len(c.BackgroundColor) != 7 && len(c.BackgroundColor) != 9 {
		errs = append(errs, errors.InvalidFieldError(constants.BackgroundColorInvalidCharsCountMsg))
	}

	if utils.IsBlank(c.TextColor) {
		errs = append(errs, errors.InvalidFieldError(constants.TextColorEmptyMsg))
	}

	if len(c.TextColor) != 7 && len(c.TextColor) != 9 {
		errs = append(errs, errors.InvalidFieldError(constants.TextColorInvalidCharsCountMsg))
	}

	return errs
}
