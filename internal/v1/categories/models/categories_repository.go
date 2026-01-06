package models

import (
	"context"
	"financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type CategoriesRepository interface {
	CreateCategory(context context.Context, data CreateCategory) (Category, []errors.ApiError)
	ReadCategoriesByUser(context context.Context, userID uuid.UUID) ([]Category, []errors.ApiError)
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (Category, []errors.ApiError)
	GetCategoriesCountByUser(context context.Context, userID uuid.UUID) (int64, []errors.ApiError)
	UpdateCategory(context context.Context, category Category) (Category, []errors.ApiError)
	DeleteCategory(context context.Context, categoryID uuid.UUID) []errors.ApiError
}
