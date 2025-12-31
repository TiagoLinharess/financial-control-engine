package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
)

func (c CategoriesService) CreateCategory(w http.ResponseWriter, r *http.Request) (categoriesModels.CategoryResponse, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	request, errs := utils.DecodeValidJson[categoriesModels.CategoryRequest](r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	data := categoriesModels.CreateCategory{
		UserID:          userID,
		TransactionType: *request.TransactionType,
		Name:            request.Name,
		Icon:            request.Icon,
	}

	category, errs := c.repository.CreateCategory(data)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return categoriesModels.CategoryResponse{
		ID:              category.ID,
		TransactionType: category.TransactionType,
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt,
		UpdatedAt:       category.UpdatedAt,
	}, http.StatusCreated, nil
}
