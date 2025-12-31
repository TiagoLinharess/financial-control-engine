package utils

import (
	"encoding/json"
	"financialcontrol/internal/models/errors"
	"net/http"
)

func DecodeJson[T any](r *http.Request) (T, errors.ApiError) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, errors.DecodeJsonError{}
	}

	return data, nil
}
