package errors

type ApiError interface {
	GetStatus() int
	GetMessages() []ApiErrorItem
}

type ApiErrorItem struct {
	UserMessage   string
	SystemMessage string
	SystemDetail  string
}

type apiError struct {
	Status   int
	Messages []ApiErrorItem
}

func NewApiErrorWithErrors(status int, messages []ApiErrorItem) ApiError {
	return apiError{
		Status:   status,
		Messages: messages,
	}
}

func NewApiError(status int, message ApiErrorItem) ApiError {
	return apiError{
		Status:   status,
		Messages: []ApiErrorItem{message},
	}
}

func (e apiError) GetStatus() int {
	return e.Status
}

func (e apiError) GetMessages() []ApiErrorItem {
	return e.Messages
}
