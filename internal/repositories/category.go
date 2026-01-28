package repositories

import (
	"context"
	"financialcontrol/internal/models"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type Category interface {
	CreateCategory(context context.Context, data models.CreateCategory) (models.Category, error)
	ReadCategories(context context.Context, userID uuid.UUID) ([]models.Category, error)
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (models.Category, error)
	GetCategoryCountByUser(context context.Context, userID uuid.UUID) (int64, error)
	UpdateCategory(context context.Context, category models.Category) (models.Category, error)
	DeleteCategory(context context.Context, categoryID uuid.UUID) error
	HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, error)
}

func (r Repository) CreateCategory(context context.Context, data models.CreateCategory) (models.Category, error) {
	param := pgstore.CreateCategoryParams{
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
	}

	category, err := r.store.CreateCategory(context, param)

	if err != nil {
		return models.Category{}, err
	}

	return StoreCategoryModelToCategory(category), nil
}

func (r Repository) ReadCategories(context context.Context, userID uuid.UUID) ([]models.Category, error) {
	rows, err := r.store.GetCategoriesByUserID(context, userID)

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return []models.Category{}, nil
	}

	result := make([]models.Category, 0, len(rows))

	for _, row := range rows {
		result = append(result, StoreCategoryModelToCategory(row))
	}

	return result, nil
}

func (r Repository) ReadCategoryByID(context context.Context, categoryID uuid.UUID) (models.Category, error) {
	category, err := r.store.GetCategoryByID(context, categoryID)

	if err != nil {
		return models.Category{}, err
	}

	return StoreCategoryModelToCategory(category), nil
}

func (r Repository) GetCategoryCountByUser(context context.Context, userID uuid.UUID) (int64, error) {
	count, err := r.store.CountCategoriesByUserID(context, userID)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r Repository) UpdateCategory(context context.Context, category models.Category) (models.Category, error) {
	param := pgstore.UpdateCategoryParams{
		ID:              category.ID,
		Name:            category.Name,
		Icon:            category.Icon,
		TransactionType: int32(category.TransactionType),
	}

	updatedCategory, err := r.store.UpdateCategory(context, param)

	if err != nil {
		return models.Category{}, err
	}

	return StoreCategoryModelToCategory(updatedCategory), nil
}

func (r Repository) DeleteCategory(context context.Context, categoryID uuid.UUID) error {
	err := r.store.DeleteCategoryByID(context, categoryID)

	if err != nil {
		return err
	}
	return nil
}

func (r Repository) HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, error) {
	hasTransactions, err := r.store.HasTransactionsByCategory(context, categoryID)

	if err != nil {
		return false, err
	}

	return hasTransactions, nil
}

func StoreCategoryModelToCategory(category pgstore.Category) models.Category {
	return models.Category{
		ID:              category.ID,
		UserID:          category.UserID,
		TransactionType: models.TransactionType(category.TransactionType),
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt.Time,
		UpdatedAt:       category.UpdatedAt.Time,
	}
}
