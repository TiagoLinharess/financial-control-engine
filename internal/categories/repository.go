package categories

import (
	"context"
	"financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type Repository interface {
	CreateCategory(context context.Context, data CreateCategory) (Category, []errors.ApiError)
	ReadCategories(context context.Context, userID uuid.UUID) ([]Category, []errors.ApiError)
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (Category, []errors.ApiError)
	GetCategoryCountByUser(context context.Context, userID uuid.UUID) (int64, []errors.ApiError)
	UpdateCategory(context context.Context, category Category) (Category, []errors.ApiError)
	DeleteCategory(context context.Context, categoryID uuid.UUID) []errors.ApiError
	HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, []errors.ApiError)
}
