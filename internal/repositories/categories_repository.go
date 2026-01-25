package repositories

import (
	"context"
	"financialcontrol/internal/categories"
	e "financialcontrol/internal/models/errors"
	"financialcontrol/internal/repositories/dtos"
	pgs "financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

func (r Repository) CreateCategory(context context.Context, data categories.CreateCategory) (categories.Category, []e.ApiError) {
	param := pgs.CreateCategoryParams{
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
	}

	category, err := r.store.CreateCategory(context, param)

	if err != nil {
		return categories.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCategoryModelToCategory(category), nil
}

func (r Repository) ReadCategories(context context.Context, userID uuid.UUID) ([]categories.Category, []e.ApiError) {
	rows, err := r.store.GetCategoriesByUserID(context, userID)

	if err != nil {
		return nil, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(rows) == 0 {
		return []categories.Category{}, nil
	}

	result := make([]categories.Category, 0, len(rows))

	for _, row := range rows {
		result = append(result, dtos.StoreCategoryModelToCategory(row))
	}

	return result, nil
}

func (r Repository) ReadCategoryByID(context context.Context, categoryID uuid.UUID) (categories.Category, []e.ApiError) {
	category, err := r.store.GetCategoryByID(context, categoryID)

	if err != nil {
		return categories.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCategoryModelToCategory(category), nil
}

func (r Repository) GetCategoryCountByUser(context context.Context, userID uuid.UUID) (int64, []e.ApiError) {
	count, err := r.store.CountCategoriesByUserID(context, userID)

	if err != nil {
		return 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return count, nil
}

func (r Repository) UpdateCategory(context context.Context, category categories.Category) (categories.Category, []e.ApiError) {
	param := pgs.UpdateCategoryParams{
		ID:              category.ID,
		Name:            category.Name,
		Icon:            category.Icon,
		TransactionType: int32(category.TransactionType),
	}

	updatedCategory, err := r.store.UpdateCategory(context, param)

	if err != nil {
		return categories.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCategoryModelToCategory(updatedCategory), nil
}

func (r Repository) DeleteCategory(context context.Context, categoryID uuid.UUID) []e.ApiError {
	err := r.store.DeleteCategoryByID(context, categoryID)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}
	return nil
}

func (r Repository) HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, []e.ApiError) {
	hasTransactions, err := r.store.HasTransactionsByCategory(context, categoryID)

	if err != nil {
		return false, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return hasTransactions, nil
}
