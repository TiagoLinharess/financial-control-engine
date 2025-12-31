package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (c CategoriesService) CreateCategory(w http.ResponseWriter, r *http.Request) (categoriesModels.CategoryResponse, errors.ErrorResponse) {
	createCategoryRequest, err := utils.DecodeJson[categoriesModels.CategoryRequest](r)

	if err != nil {
		return categoriesModels.CategoryResponse{}, errors.NewErrorResponse(http.StatusUnprocessableEntity, []errors.ApiError{err})
	}

	return categoriesModels.CategoryResponse{
		ID:              uuid.UUID{},
		TransactionType: createCategoryRequest.TransactionType,
		Name:            createCategoryRequest.Name,
		Icon:            createCategoryRequest.Icon,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, errors.EmptyErrorResponse()
}
