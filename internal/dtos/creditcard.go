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

func (c CreditCardRequest) Validate() []errors.ApiError {
	errs := make([]errors.ApiError, 0)

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardNameEmptyMsg})
	}

	if len(c.Name) >= 100 || len(c.Name) <= 2 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardNameInvalidCharsCountMsg})
	}

	if len(c.FirstFourNumbers) != 4 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardFirstFourNumbersInvalidMsg})
	}

	if c.Limit <= 0 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardLimitInvalidMsg})
	}

	if c.CloseDay < 1 || c.CloseDay > 31 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardClosingDayInvalidMsg})
	}

	if c.ExpireDay < 1 || c.ExpireDay > 31 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardExpireDayInvalidMsg})
	}

	if utils.IsBlank(c.BackgroundColor) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardBackgroundColorEmptyMsg})
	}

	if len(c.BackgroundColor) != 7 && len(c.BackgroundColor) != 9 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardBackgroundColorInvalidCharsCountMsg})
	}

	if utils.IsBlank(c.TextColor) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardTextColorEmptyMsg})
	}

	if len(c.TextColor) != 7 && len(c.TextColor) != 9 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CreditcardTextColorInvalidCharsCountMsg})
	}

	return errs
}
