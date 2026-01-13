package errors

type CustomError struct {
	Message string
}

func (c CustomError) String() string {
	return c.Message
}
