package models

import (
	"context"
	"financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type CategoriesRepository interface {
	Create(context context.Context, data CreateCategory) (Category, []errors.ApiError)
	Read(context context.Context, userID uuid.UUID) ([]Category, []errors.ApiError)
	ReadByID(context context.Context, categoryID uuid.UUID) (Category, []errors.ApiError)
	GetCountByUser(context context.Context, userID uuid.UUID) (int64, []errors.ApiError)
	Update(context context.Context, category Category) (Category, []errors.ApiError)
	Delete(context context.Context, categoryID uuid.UUID) []errors.ApiError
}
