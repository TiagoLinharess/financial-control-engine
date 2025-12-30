package services

import (
	models "financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (c CategoriesService) CreateCategory(w http.ResponseWriter, r *http.Request) (categoriesModels.CategoryResponse, errors.ErrorResponse) {
	return categoriesModels.CategoryResponse{
		ID:              uuid.UUID{},
		TransactionType: models.Income,
		Name:            "name",
		Icon:            "add",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, errors.EmptyErrorResponse()
}
