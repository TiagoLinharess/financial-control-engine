package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"
)

type CategoriesStore interface {
	CreateCategory(ctx context.Context, arg pgstore.CreateCategoryParams) (pgstore.Category, error)
}
