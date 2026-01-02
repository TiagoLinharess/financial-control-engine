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
}
