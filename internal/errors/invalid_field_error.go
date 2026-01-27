package errors

import (
	"financialcontrol/internal/constants"
)

type InvalidFieldError struct {
	Message string
}

func (d InvalidFieldError) String() string {
	return d.Message
}

func (d InvalidFieldError) SystemMessage() (string, string) {
	return constants.InvalidFieldErrorSystemMsg, d.Message
}
