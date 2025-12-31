package models

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
)

type CategoryRequest struct {
	TransactionType *models.TransactionType `json:"transaction_type"`
	Name            string                  `json:"name"`
	Icon            string                  `json:"icon"`
}

func (c CategoryRequest) Validate() []errors.ApiError {
	errs := make([]errors.ApiError, 0)
	if c.TransactionType == nil {
		errs = append(errs, errors.InvalidFieldError{Message: "Transaction id must not be empty"})
	} else if !c.TransactionType.IsValid() {
		errs = append(errs, errors.InvalidFieldError{Message: "Transaction id must be valid"})
	}

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: "Name musto not be empty"})
	}

	if utils.IsBlank(c.Icon) {
		errs = append(errs, errors.InvalidFieldError{Message: "Icon musto not be empty"})
	}

	return errs
}
