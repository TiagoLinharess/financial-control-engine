package errors

type DecodeJsonError struct {
}

func (d DecodeJsonError) String() string {
	return "Error on deconding json"
}
