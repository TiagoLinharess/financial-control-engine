package errors

type UnauthorizedErrorReasons string

const (
	UserIDNotFound UnauthorizedErrorReasons = "UserID not found"
	UserIDInvalid  UnauthorizedErrorReasons = "UserID invalid"
)

type UnauthorizedError struct {
	Message UnauthorizedErrorReasons
}

func (u UnauthorizedError) String() string {
	return string(u.Message)
}
