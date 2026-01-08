package errors

type LimitErrorReasons string

const (
	CategoriesLimit  LimitErrorReasons = "You reached the maximum number of categories allowed"
	CreditcardsLimit LimitErrorReasons = "You reached the maximum number of credit cards allowed"
)

type LimitError struct {
	Message LimitErrorReasons
}

func (l LimitError) String() string {
	return string(l.Message)
}
