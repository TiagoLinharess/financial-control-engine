package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s CategoriesService) DeleteCategory(ctx *gin.Context) (int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param("id")

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(ctx, categoryID)

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

	errs = s.repository.DeleteCategory(ctx, categoryID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}
