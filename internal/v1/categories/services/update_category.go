package services

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s CategoriesService) UpdateCategory(ctx *gin.Context) (categoriesModels.CategoryResponse, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param("id")

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(ctx, categoryID)

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

	request, errs := utils.DecodeValidJson[categoriesModels.CategoryRequest](ctx)

	if len(errs) > 0 {
		return categoriesModels.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, errs := s.repository.UpdateCategory(ctx, category)

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
