package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	"financialcontrol/internal/v1/categories/models"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s CategoriesService) ReadCategoryByID(w http.ResponseWriter, r *http.Request) (models.CategoryResponse, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	categoryIDString := chi.URLParam(r, "id")

	categoryID, err := uuid.Parse(categoryIDString)
	if err != nil {
		return models.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(r.Context(), categoryID)

	if len(errs) > 0 || errs.cont {
		return models.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID || category.ID != categoryID {
		return models.CategoryResponse{}, http.StatusNotFound, []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	}

	return models.CategoryResponse{
		ID:              category.ID,
		TransactionType: category.TransactionType,
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt,
		UpdatedAt:       category.UpdatedAt,
	}, http.StatusOK, nil
}
