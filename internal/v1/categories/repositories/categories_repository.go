package repositories

import (
	"context"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store"
	pgs "financialcontrol/internal/store/pgstore"
	cm "financialcontrol/internal/v1/categories/models"

	"github.com/google/uuid"
)

type CategoriesRepository struct {
	store st.CategoriesStore
}

func NewCategoriesRepository(store st.CategoriesStore) cm.CategoriesRepository {
	return CategoriesRepository{store: store}
}

func (c CategoriesRepository) Create(context context.Context, data cm.CreateCategory) (cm.Category, []e.ApiError) {
	param := pgs.CreateCategoryParams{
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
	}

	category, err := c.store.CreateCategory(context, param)

	if err != nil {
		return cm.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(category), nil
}

func (c CategoriesRepository) Read(context context.Context, userID uuid.UUID) ([]cm.Category, []e.ApiError) {
	categories, err := c.store.GetCategoriesByUserID(context, userID)

	if err != nil {
		return nil, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(categories) == 0 {
		return []cm.Category{}, nil
	}

	categoriesResponse := make([]cm.Category, 0, len(categories))

	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, storeModelToModel(category))
	}

	return categoriesResponse, nil
}

func (c CategoriesRepository) ReadByID(context context.Context, categoryID uuid.UUID) (cm.Category, []e.ApiError) {
	category, err := c.store.GetCategoryByID(context, categoryID)

	if err != nil {
		return cm.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(category), nil
}

func (c CategoriesRepository) GetCountByUser(context context.Context, userID uuid.UUID) (int64, []e.ApiError) {
	count, err := c.store.CountCategoriesByUserID(context, userID)

	if err != nil {
		return 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return count, nil
}

func (c CategoriesRepository) Update(context context.Context, category cm.Category) (cm.Category, []e.ApiError) {
	param := pgs.UpdateCategoryParams{
		ID:              category.ID,
		Name:            category.Name,
		Icon:            category.Icon,
		TransactionType: int32(category.TransactionType),
	}

	updatedCategory, err := c.store.UpdateCategory(context, param)

	if err != nil {
		return cm.Category{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(updatedCategory), nil
}

func (c CategoriesRepository) Delete(context context.Context, categoryID uuid.UUID) []e.ApiError {
	err := c.store.DeleteCategoryByID(context, categoryID)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}
	return nil
}

func (c CategoriesRepository) HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, []e.ApiError) {
	hasTransactions, err := c.store.HasTransactionsByCategory(context, categoryID)

	if err != nil {
		return false, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return hasTransactions, nil
}

func storeModelToModel(category pgs.Category) cm.Category {
	return cm.Category{
		ID:              category.ID,
		UserID:          category.UserID,
		TransactionType: m.TransactionType(category.TransactionType),
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt.Time,
		UpdatedAt:       category.UpdatedAt.Time,
	}
}
