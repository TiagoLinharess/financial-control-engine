package errors

import (
	"financialcontrol/internal/constants"
)

type LimitErrorReasons string

const (
	CategoriesLimit  LimitErrorReasons = constants.CategoryLimitReachedMsg
	CreditcardsLimit LimitErrorReasons = constants.CreditcardLimitReachedMsg
)

type LimitError struct {
	Message LimitErrorReasons
}

func (l LimitError) String() string {
	return string(l.Message)
}

func (l LimitError) SystemMessage() (string, string) {
	return constants.LimitErrorSystemMsg, string(l.Message)
}
