package utils

import (
	"encoding/json"
	"financialcontrol/internal/models/errors"
	"net/http"
)

func SendResponse[T any](w http.ResponseWriter, object T, status int) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(object); err != nil {
		SendResponse(w, "error", http.StatusInternalServerError)
		return
	}
}

func SendError(w http.ResponseWriter, error errors.ErrorResponse, status int) {
	SendResponse(w, error, status)
}
