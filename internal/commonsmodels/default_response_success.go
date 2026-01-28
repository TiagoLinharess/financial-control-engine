package commonsmodels

import "financialcontrol/internal/constants"

type ResponseSuccess struct {
	Message string `json:"message"`
}

func NewResponseSuccess() ResponseSuccess {
	return ResponseSuccess{Message: constants.Success}
}
