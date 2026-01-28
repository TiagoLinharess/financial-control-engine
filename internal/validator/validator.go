package validator

import "financialcontrol/internal/errors"

type Validator interface {
	Validate() []errors.ApiErrorItem
}
