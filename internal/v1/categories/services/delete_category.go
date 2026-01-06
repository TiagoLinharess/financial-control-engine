package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (s CategoriesService) DeleteCategory(w http.ResponseWriter, r *http.Request) (int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return http.StatusUnauthorized, errs
	}

	categoryIDString := chi.URLParam(r, "id")

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(r.Context(), categoryID)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == string(store.ErrNoRows)
		})
		if isNotFoundErr {
			return http.StatusNotFound, categoryNotFoundErr
		}
		return http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return http.StatusNotFound, categoryNotFoundErr
	}

	errs = s.repository.DeleteCategory(r.Context(), categoryID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}
