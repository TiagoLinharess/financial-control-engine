package errors

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func NewErrorResponse(errors []ApiError) ErrorResponse {
	errorMessages := make([]string, len(errors))

	for i, p := range errors {
		errorMessages[i] = p.String()
	}

	return ErrorResponse{Errors: errorMessages}
}
