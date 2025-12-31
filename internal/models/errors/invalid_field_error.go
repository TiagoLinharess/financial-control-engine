package errors

type InvalidFieldError struct {
	Message string
}

func (d InvalidFieldError) String() string {
	return d.Message
}
