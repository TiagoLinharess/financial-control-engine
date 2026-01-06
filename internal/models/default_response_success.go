package models

type ResponseSuccess struct {
	Message string `json:"message"`
}

func NewResponseSuccess() ResponseSuccess {
	return ResponseSuccess{Message: "success"}
}
