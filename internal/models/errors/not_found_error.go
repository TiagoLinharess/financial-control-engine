package errors

type NotFoundErrorType string

const (
	CategoryNotFound NotFoundErrorType = "Category not found"
)

type NotFoundError struct {
	Message NotFoundErrorType
}

func (n NotFoundError) String() string {
	return string(n.Message)
}
