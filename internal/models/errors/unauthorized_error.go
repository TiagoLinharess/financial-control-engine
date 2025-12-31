package errors

type UnauthorizedErrorReasons string

const (
	UserIDNotFound = "UserID not found"
	UserIDInvalid  = "UserID invalid"
)

type UnauthorizedError struct {
	Message string
}

func (u UnauthorizedError) String() string {
	return u.Message
}
