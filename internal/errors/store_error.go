package errors

import "financialcontrol/internal/constants"

type StoreError struct {
	Message string
}

func (s StoreError) String() string {
	return s.Message
}

func (s StoreError) SystemMessage() (string, string) {
	return constants.StoreErrorSystemMsg, s.Message
}
