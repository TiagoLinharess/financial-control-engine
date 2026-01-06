package repositories

import (
	"context"
	models "financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/store"
	"financialcontrol/internal/store/pgstore"
	categoriesModels "financialcontrol/internal/v1/categories/models"

	"github.com/google/uuid"
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

func (c CategoriesRepository) ReadCategoriesByUser(context context.Context, userID uuid.UUID) ([]categoriesModels.Category, []errors.ApiError) {
	categories, err := c.store.GetCategoriesByUserID(context, userID)

	if err != nil {
		return nil, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	if len(categories) == 0 {
		return []categoriesModels.Category{}, nil
	}

	categoriesResponse := make([]categoriesModels.Category, 0, len(categories))

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, categoriesModels.Category{
			ID:              category.ID,
			UserID:          category.UserID,
			TransactionType: models.TransactionType(category.TransactionType),
			Name:            category.Name,
			Icon:            category.Icon,
			CreatedAt:       category.CreatedAt.Time,
			UpdatedAt:       category.UpdatedAt.Time,
		})
	}

	return categoriesResponse, nil
}

func (c CategoriesRepository) ReadCategoryByID(context context.Context, categoryID uuid.UUID) (categoriesModels.Category, []errors.ApiError) {
	category, err := c.store.GetCategoryByID(context, categoryID)

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

func (c CategoriesRepository) GetCategoriesCountByUser(context context.Context, userID uuid.UUID) (int64, []errors.ApiError) {
	count, err := c.store.CountCategoriesByUserID(context, userID)

	if err != nil {
		return 0, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return count, nil
}

func (c CategoriesRepository) UpdateCategory(context context.Context, category categoriesModels.Category) (categoriesModels.Category, []errors.ApiError) {
	param := pgstore.UpdateCategoryParams{
		ID:              category.ID,
		Name:            category.Name,
		Icon:            category.Icon,
		TransactionType: int32(category.TransactionType),
	}

	updatedCategory, err := c.store.UpdateCategory(context, param)

	if err != nil {
		return categoriesModels.Category{}, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return categoriesModels.Category{
		ID:              updatedCategory.ID,
		UserID:          updatedCategory.UserID,
		TransactionType: models.TransactionType(updatedCategory.TransactionType),
		Name:            updatedCategory.Name,
		Icon:            updatedCategory.Icon,
		CreatedAt:       updatedCategory.CreatedAt.Time,
		UpdatedAt:       updatedCategory.UpdatedAt.Time,
	}, nil
}

func (c CategoriesRepository) DeleteCategory(context context.Context, categoryID uuid.UUID) []errors.ApiError {
	err := c.store.DeleteCategoryByID(context, categoryID)

	if err != nil {
		return []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}
	return nil
}
