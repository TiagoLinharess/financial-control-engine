package utils

import (
	"encoding/json"
	"financialcontrol/internal/models/errors"
	"log"
	"net/http"
)

func SendResponse[T any](w http.ResponseWriter, data T, status int, errs []errors.ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var response interface{}

	if len(errs) > 0 {
		response = errors.NewErrorResponse(errs)
	} else {
		response = data
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		return
	}
}
