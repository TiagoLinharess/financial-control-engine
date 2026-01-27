package utils

import (
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"

	"github.com/gin-gonic/gin"
)

func DecodeJson[T any](ctx *gin.Context) (T, []errors.ApiError) {
	var data T

	if err := ctx.BindJSON(&data); err != nil {
		return data, []errors.ApiError{errors.DecodeJsonError{}}
	}

	return data, nil
}

func DecodeValidJson[T models.Validator](ctx *gin.Context) (T, []errors.ApiError) {
	data, errs := DecodeJson[T](ctx)

	if errs != nil {
		return data, errs
	}

	errs = data.Validate()

	return data, errs
}
