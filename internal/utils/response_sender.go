package utils

import (
	"encoding/json"
	"financialcontrol/internal/models/errors"
	errorsModel "financialcontrol/internal/models/errors"
	"net/http"
)

func SendResponse[T any](w http.ResponseWriter, data T, status int, errors []errors.ApiError) {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(status)

	if len(errors) > 0 {
		sendJson(
			w,
			errorsModel.NewErrorResponse(errors),
		)
		return
	}

	sendJson(w, data)
}

func sendJson[T any](w http.ResponseWriter, data T) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		sendJson(
			w,
			errorsModel.NewErrorResponse([]errorsModel.ApiError{errorsModel.EncodeJsonError{}}),
		)
		return
	}
}
