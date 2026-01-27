package models

import "financialcontrol/internal/errors"

type Validator interface {
	Validate() []errors.ApiError
}
