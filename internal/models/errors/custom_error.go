package errors

import (
	"financialcontrol/internal/constants"
)

type CustomError struct {
	Message string
}

func (c CustomError) String() string {
	return c.Message
}

func (c CustomError) SystemMessage() (string, string) {
	return constants.CustomError, c.Message
}
