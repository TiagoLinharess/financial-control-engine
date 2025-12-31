package errors

type EncodeJsonError struct {
}

func (d EncodeJsonError) String() string {
	return "Error on enconding json"
}
