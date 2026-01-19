package errors

import "financialcontrol/internal/constants"

type EncodeJsonError struct {
}

func (d EncodeJsonError) String() string {
	return constants.EncodeJsonErrorMsg
}

func (d EncodeJsonError) SystemMessage() (string, string) {
	return constants.EncodeJsonErrorSystemMsg, constants.EmptyString
}
