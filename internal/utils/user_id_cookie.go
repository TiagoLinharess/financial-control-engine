package utils

import (
	"financialcontrol/internal/models/errors"
	"net/http"

	"github.com/google/uuid"
)

func ReadUserIdFromCookie(w http.ResponseWriter, r *http.Request) (uuid.UUID, []errors.ApiError) {
	cookie, err := r.Cookie("user_id")

	if err != nil {
		return uuid.UUID{}, []errors.ApiError{errors.UnauthorizedError{Message: errors.UserIDNotFound}}
	}

	userIDString := cookie.Value
	userID, err := uuid.Parse(userIDString)

	if err != nil {
		return uuid.UUID{}, []errors.ApiError{errors.UnauthorizedError{Message: errors.UserIDInvalid}}
	}

	return userID, nil
}
