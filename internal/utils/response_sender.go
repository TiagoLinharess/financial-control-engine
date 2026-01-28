package utils

import (
	"financialcontrol/internal/errors"

	"github.com/gin-gonic/gin"
)

func SendResponse[T any](ctx *gin.Context, data T, status int) {
	// TODO: LOCAL FILE LOGS
	ctx.JSON(status, data)
}

func SendErrorResponse(ctx *gin.Context, err errors.ApiError) {
	var userMessages []string
	for _, msg := range err.GetMessages() {
		userMessages = append(userMessages, msg.UserMessage)
	}

	response := errors.ErrorResponse{
		Status:   err.GetStatus(),
		Messages: userMessages,
	}

	// TODO: LOCAL FILE LOGS
	ctx.AbortWithStatusJSON(err.GetStatus(), response)
}
