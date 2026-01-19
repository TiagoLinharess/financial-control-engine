package errors

import "financialcontrol/internal/constants"

type DecodeJsonError struct {
}

func (d DecodeJsonError) String() string {
	return constants.DecodeJsonErrorMsg
}

func (d DecodeJsonError) SystemMessage() (string, string) {
	return constants.DecodeJsonErrorSystemMsg, constants.EmptyString
}
