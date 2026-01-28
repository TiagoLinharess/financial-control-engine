package errors

type ErrorResponse struct {
	Status   int      `json:"status"`
	Messages []string `json:"messages"`
}
