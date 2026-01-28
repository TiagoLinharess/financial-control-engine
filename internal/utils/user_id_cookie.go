package utils

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIDFromContext(ctx *gin.Context) (uuid.UUID, errors.ApiError) {
	userID, exists := ctx.Get(constants.UserID)
	if !exists {
		return uuid.UUID{}, errors.NewApiErrorWithErrors(http.StatusUnauthorized, []errors.ApiErrorItem{errors.UserIDInvalid("")})
	}

	id, ok := userID.(uuid.UUID)

	if !ok {
		return uuid.UUID{}, errors.NewApiErrorWithErrors(http.StatusUnauthorized, []errors.ApiErrorItem{errors.UserIDInvalid("")})
	}

	return id, nil
}
