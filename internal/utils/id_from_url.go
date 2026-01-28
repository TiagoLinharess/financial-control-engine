package utils

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IDFromURLParam(ctx *gin.Context) (uuid.UUID, errors.ApiError) {
	idString := ctx.Param(constants.ID)

	id, err := uuid.Parse(idString)

	if err != nil {
		return uuid.UUID{}, errors.NewApiError(http.StatusBadRequest, errors.BadRequestError(constants.InvalidID))
	}

	return id, nil
}
