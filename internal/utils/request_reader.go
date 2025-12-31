package utils

import (
	"encoding/json"
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"net/http"
)

func DecodeJson[T any](r *http.Request) (T, []errors.ApiError) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, []errors.ApiError{errors.DecodeJsonError{}}
	}

	return data, nil
}

func DecodeValidJson[T models.Validator](r *http.Request) (T, []errors.ApiError) {
	data, errs := DecodeJson[T](r)

	if errs != nil {
		return data, errs
	}

	errs = data.Validate()

	return data, errs
}
