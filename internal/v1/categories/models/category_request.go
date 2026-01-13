package models

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"

	"github.com/google/uuid"
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
		errs = append(errs, errors.InvalidFieldError{Message: "Name must not be empty"})
	}

	if utils.IsBlank(c.Icon) {
		errs = append(errs, errors.InvalidFieldError{Message: "Icon must not be empty"})
	}

	return errs
}

func (c CategoryRequest) ToCreateModel(userID uuid.UUID) CreateCategory {
	return CreateCategory{
		UserID:          userID,
		TransactionType: *c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}
