package services

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c CategoriesService) Read(ctx *gin.Context) (m.ResponseList[cm.CategoryResponse], int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return m.ResponseList[cm.CategoryResponse]{}, http.StatusUnauthorized, errs
	}

	categories, errs := c.repository.Read(ctx, userID)

	if len(errs) > 0 {
		return m.ResponseList[cm.CategoryResponse]{}, http.StatusInternalServerError, errs
	}

	response := make([]cm.CategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, category.ToResponse())
	}

	return m.ResponseList[cm.CategoryResponse]{Items: response, Total: len(response)}, http.StatusOK, nil
}
