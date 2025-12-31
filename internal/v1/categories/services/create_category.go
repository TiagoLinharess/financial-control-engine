package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (c CategoriesService) CreateCategory(w http.ResponseWriter, r *http.Request) (categoriesModels.CategoryResponse, int, []errors.ApiError) {
	request, errs := utils.DecodeValidJson[categoriesModels.CategoryRequest](r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	return categoriesModels.CategoryResponse{
		ID:              uuid.UUID{},
		TransactionType: *request.TransactionType,
		Name:            request.Name,
		Icon:            request.Icon,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, http.StatusCreated, nil
}
