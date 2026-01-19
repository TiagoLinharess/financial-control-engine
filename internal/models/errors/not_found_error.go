package errors

import "financialcontrol/internal/constants"

type NotFoundErrorType string

const (
	CategoryNotFound    NotFoundErrorType = constants.CategoryNotFoundMsg
	CreditcardNotFound  NotFoundErrorType = constants.CreditcardNotFoundMsg
	TransactionNotFound NotFoundErrorType = constants.TransactionNotFoundMsg
)

type NotFoundError struct {
	Message NotFoundErrorType
}

func (n NotFoundError) String() string {
	return string(n.Message)
}

func (n NotFoundError) SystemMessage() (string, string) {
	return constants.NotFoundErrorSystemMsg, string(n.Message)
}
