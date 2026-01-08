package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CategoriesService) Create(ctx *gin.Context) (cm.CategoryResponse, int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.GetCountByUser(ctx, userID)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return cm.CategoryResponse{}, http.StatusForbidden, []e.ApiError{e.LimitError{Message: e.CategoriesLimit}}
	}

	request, errs := u.DecodeValidJson[cm.CategoryRequest](ctx)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusBadRequest, errs
	}

	data := request.ToCreateModel(userID)

	category, errs := c.repository.Create(ctx, data)

	if len(errs) > 0 {
		return cm.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return category.ToResponse(), http.StatusCreated, nil
}
