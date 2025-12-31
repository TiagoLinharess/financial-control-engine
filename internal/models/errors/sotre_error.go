package errors

type StoreError struct {
	Message string
}

func (s StoreError) String() string {
	return s.Message
}
