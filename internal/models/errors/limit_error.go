package errors

type LimitErrorReasons string

const (
	CategoriesLimit LimitErrorReasons = "You reached the maximum number of categories allowed"
)

type LimitError struct {
	Message LimitErrorReasons
}

func (l LimitError) String() string {
	return string(l.Message)
}
