package services

import (
	e "financialcontrol/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s CategoriesService) Delete(ctx *gin.Context) (int, []e.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return statusCode, errs
	}

	// TODO: check if category has transactions associated

	errs = s.repository.Delete(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}
