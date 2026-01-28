package middlewares

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UserIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDString, err := ctx.Cookie(constants.UserID)

		if err != nil {
			utils.SendErrorResponse(
				ctx,
				errors.NewApiError(http.StatusUnauthorized, errors.UserNotFound(err.Error())),
			)
			return
		}

		userID, err := uuid.Parse(userIDString)

		if err != nil {
			utils.SendErrorResponse(
				ctx,
				errors.NewApiError(http.StatusUnauthorized, errors.UserNotFound(err.Error())),
			)
			return
		}

		ctx.Set(constants.UserID, userID)
		ctx.Next()
	}
}
