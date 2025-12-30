package errors

import "net/http"

type ErrorResponse struct {
	Status int      `json:"-"`
	Errors []string `json:"errors"`
}

func NewErrorResponse(status int, errors []ApiError) ErrorResponse {
	errorMessages := make([]string, len(errors))

	for i, p := range errors {
		errorMessages[i] = p.String()
	}

	return ErrorResponse{Status: status, Errors: errorMessages}
}

func EmptyErrorResponse() ErrorResponse {
	return ErrorResponse{Status: http.StatusInternalServerError}
}
