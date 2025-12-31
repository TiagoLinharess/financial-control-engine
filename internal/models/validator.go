package models

import "financialcontrol/internal/models/errors"

type Validator interface {
	Validate() []errors.ApiError
}
