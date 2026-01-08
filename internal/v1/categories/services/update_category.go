package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s CategoriesService) Update(ctx *gin.Context) (cm.CategoryResponse, int, []e.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, statusCode, errs
	}

	request, errs := u.DecodeValidJson[cm.CategoryRequest](ctx)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, errs := s.repository.Update(ctx, category)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return categoryEdited.ToResponse(), http.StatusOK, nil
}
