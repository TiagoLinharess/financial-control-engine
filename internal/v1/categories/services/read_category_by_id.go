package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/utils"
	"financialcontrol/internal/v1/categories/models"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s CategoriesService) ReadCategoryByID(ctx *gin.Context) (models.CategoryResponse, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param("id")

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return models.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(ctx, categoryID)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == string(store.ErrNoRows)
		})
		if isNotFoundErr {
			return models.CategoryResponse{}, http.StatusNotFound, categoryNotFoundErr
		}
		return models.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return models.CategoryResponse{}, http.StatusNotFound, categoryNotFoundErr
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
