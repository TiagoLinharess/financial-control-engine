package errors

type ApiError interface {
	String() string
	SystemMessage() (string, string)
}
