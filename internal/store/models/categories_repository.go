package models

import (
	"context"
	e "financialcontrol/internal/models/errors"
	cm "financialcontrol/internal/v1/categories/models"

	"github.com/google/uuid"
)

type CategoriesRepository interface {
	CreateCategory(context context.Context, data cm.CreateCategory) (cm.Category, []e.ApiError)
	ReadCategories(context context.Context, userID uuid.UUID) ([]cm.Category, []e.ApiError)
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (cm.Category, []e.ApiError)
	GetCategoryCountByUser(context context.Context, userID uuid.UUID) (int64, []e.ApiError)
	UpdateCategory(context context.Context, category cm.Category) (cm.Category, []e.ApiError)
	DeleteCategory(context context.Context, categoryID uuid.UUID) []e.ApiError
	HasTransactionsByCategory(context context.Context, categoryID uuid.UUID) (bool, []e.ApiError)
}
