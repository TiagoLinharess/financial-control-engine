package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type CategoriesStore interface {
	CreateCategory(ctx context.Context, arg pgstore.CreateCategoryParams) (pgstore.Category, error)
	GetCategoriesByUserID(ctx context.Context, userID uuid.UUID) ([]pgstore.Category, error)
	GetCategoryByID(ctx context.Context, id uuid.UUID) (pgstore.Category, error)
	CountCategoriesByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
	UpdateCategory(ctx context.Context, arg pgstore.UpdateCategoryParams) (pgstore.Category, error)
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
	HasTransactionsByCategory(ctx context.Context, categoryID uuid.UUID) (bool, error)
}
