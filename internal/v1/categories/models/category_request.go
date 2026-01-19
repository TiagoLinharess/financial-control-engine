package models

import (
	"financialcontrol/internal/constants"
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
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryTransactionTypeEmptyMsg})
	} else if !c.TransactionType.IsValid() {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryTransactionTypeMsg})
	}

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryNameEmptyMsg})
	}

	if len(c.Name) < 2 || len(c.Name) > 255 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryNameInvalidCharsCountMsg})
	}

	if len(c.Icon) < 2 || len(c.Icon) > 255 {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryIconInvalidCharsCountMsg})
	}

	if utils.IsBlank(c.Icon) {
		errs = append(errs, errors.InvalidFieldError{Message: constants.CategoryIconEmptyMsg})
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
