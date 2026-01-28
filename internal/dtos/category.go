package dtos

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
)

type ShortCategoryResponse struct {
	ID              uuid.UUID              `json:"id"`
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
}

type CategoryResponse struct {
	ID              uuid.UUID              `json:"id"`
	TransactionType models.TransactionType `json:"transaction_type"`
	Name            string                 `json:"name"`
	Icon            string                 `json:"icon"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type CategoryRequest struct {
	TransactionType *models.TransactionType `json:"transaction_type"`
	Name            string                  `json:"name"`
	Icon            string                  `json:"icon"`
}

func (c CategoryRequest) Validate() []errors.ApiErrorItem {
	errs := make([]errors.ApiErrorItem, 0)
	if c.TransactionType == nil {
		errs = append(errs, errors.InvalidFieldError(constants.TransactionTypeEmptyMsg))
	} else if !c.TransactionType.IsValid() {
		errs = append(errs, errors.InvalidFieldError(constants.TransactionTypeMsg))
	}

	if utils.IsBlank(c.Name) {
		errs = append(errs, errors.InvalidFieldError(constants.NameEmptyMsg))
	}

	if len(c.Name) < 2 || len(c.Name) > 255 {
		errs = append(errs, errors.InvalidFieldError(constants.NameInvalidCharsCountMsg))
	}

	if len(c.Icon) < 2 || len(c.Icon) > 255 {
		errs = append(errs, errors.InvalidFieldError(constants.IconInvalidCharsCountMsg))
	}

	if utils.IsBlank(c.Icon) {
		errs = append(errs, errors.InvalidFieldError(constants.IconEmptyMsg))
	}

	return errs
}
