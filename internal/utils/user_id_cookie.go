package utils

import (
	"financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ReadUserIdFromCookie(ctx *gin.Context) (uuid.UUID, []errors.ApiError) {
	userIDString, err := ctx.Cookie("user_id")

	if err != nil {
		return uuid.UUID{}, []errors.ApiError{errors.UnauthorizedError{Message: errors.UserIDNotFound}}
	}

	userID, err := uuid.Parse(userIDString)

	if err != nil {
		return uuid.UUID{}, []errors.ApiError{errors.UnauthorizedError{Message: errors.UserIDInvalid}}
	}

	return userID, nil
}
