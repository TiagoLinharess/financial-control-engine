package models

import (
	"context"
	"financialcontrol/internal/models/errors"
)

type CategoriesRepository interface {
	CreateCategory(context context.Context, data CreateCategory) (Category, []errors.ApiError)
}
