package services

import (
	globalModels "financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
)

func (c CategoriesService) ReadCategoriesByUser(w http.ResponseWriter, r *http.Request) (globalModels.ResponseList[categoriesModels.CategoryResponse], int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(w, r)

	if len(errs) > 0 {
		return globalModels.ResponseList[categoriesModels.CategoryResponse]{}, http.StatusUnauthorized, errs
	}

	categories, errs := c.repository.ReadCategoriesByUser(r.Context(), userID)

	if len(errs) > 0 {
		return globalModels.ResponseList[categoriesModels.CategoryResponse]{}, http.StatusInternalServerError, errs
	}

	response := make([]categoriesModels.CategoryResponse, 0, len(categories))

	for _, category := range categories {
		response = append(response, categoriesModels.CategoryResponse{
			ID:              category.ID,
			TransactionType: category.TransactionType,
			Name:            category.Name,
			Icon:            category.Icon,
			CreatedAt:       category.CreatedAt,
			UpdatedAt:       category.UpdatedAt,
		})
	}

	return globalModels.ResponseList[categoriesModels.CategoryResponse]{Items: response, Total: len(response)}, http.StatusOK, nil
}
