package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s CategoriesService) UpdateCategory(w http.ResponseWriter, r *http.Request) (categoriesModels.CategoryResponse, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	categoryIDString := chi.URLParam(r, "id")

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(r.Context(), categoryID)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == string(store.ErrNoRows)
		})
		if isNotFoundErr {
			return categoriesModels.CategoryResponse{}, http.StatusNotFound, categoryNotFoundErr
		}
		return categoriesModels.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return categoriesModels.CategoryResponse{}, http.StatusNotFound, categoryNotFoundErr
	}

	request, errs := utils.DecodeValidJson[categoriesModels.CategoryRequest](r)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, errs := s.repository.UpdateCategory(r.Context(), category)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return categoriesModels.CategoryResponse{
		ID:              categoryEdited.ID,
		TransactionType: categoryEdited.TransactionType,
		Name:            categoryEdited.Name,
		Icon:            categoryEdited.Icon,
		CreatedAt:       categoryEdited.CreatedAt,
		UpdatedAt:       categoryEdited.UpdatedAt,
	}, http.StatusOK, nil
}
