package repositories

import (
	"context"
	models "financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/store/pgstore"
	categoriesModels "financialcontrol/internal/v1/categories/models"
)

type CategoriesRepository struct {
	store store.CategoriesStore
}

func NewCategoriesRepository(store store.CategoriesStore) categoriesModels.CategoriesRepository {
	return CategoriesRepository{store: store}
}

func (c CategoriesRepository) CreateCategory(context context.Context, data categoriesModels.CreateCategory) (categoriesModels.Category, []errors.ApiError) {
	param := pgstore.CreateCategoryParams{
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
	}

	category, err := c.store.CreateCategory(context, param)

	if err != nil {
		return categoriesModels.Category{}, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return categoriesModels.Category{
		ID:              category.ID,
		UserID:          category.UserID,
		TransactionType: models.TransactionType(category.TransactionType),
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt.Time,
		UpdatedAt:       category.UpdatedAt.Time,
	}, nil
}
