package models

import "financialcontrol/internal/models/errors"

type CategoriesRepository interface {
	CreateCategory(data CreateCategory) (Category, []errors.ApiError)
}
