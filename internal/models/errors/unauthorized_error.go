package errors

import constants "financialcontrol/internal/constants"

type UnauthorizedErrorReasons string

const (
	UserIDNotFound UnauthorizedErrorReasons = constants.UserIDNotFoundMsg
	UserIDInvalid  UnauthorizedErrorReasons = constants.UserIDInvalidMsg
)

type UnauthorizedError struct {
	Message UnauthorizedErrorReasons
}

func (u UnauthorizedError) String() string {
	return string(u.Message)
}

func (u UnauthorizedError) SystemMessage() (string, string) {
	return constants.UnauthorizedErrorSystemMsg, string(u.Message)
}
