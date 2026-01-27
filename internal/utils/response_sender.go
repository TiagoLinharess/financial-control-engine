package utils

import (
	"financialcontrol/internal/errors"

	"github.com/gin-gonic/gin"
)

func SendResponse[T any](ctx *gin.Context, data T, status int, errs []errors.ApiError) {
	var response interface{}

	if len(errs) > 0 {
		response = errors.NewErrorResponse(errs)
	} else {
		response = data
	}

	ctx.JSON(status, response)
}
