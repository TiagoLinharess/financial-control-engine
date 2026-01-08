package services

import (
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s CategoriesService) ReadByID(ctx *gin.Context) (cm.CategoryResponse, int, []e.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, statusCode, errs
	}

	return category.ToResponse(), http.StatusOK, nil
}
