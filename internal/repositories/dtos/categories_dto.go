package dtos

import (
	"financialcontrol/internal/categories"
	m "financialcontrol/internal/models"
	pgs "financialcontrol/internal/store/pgstore"
)

func StoreCategoryModelToCategory(category pgs.Category) categories.Category {
	return categories.Category{
		ID:              category.ID,
		UserID:          category.UserID,
		TransactionType: m.TransactionType(category.TransactionType),
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt.Time,
		UpdatedAt:       category.UpdatedAt.Time,
	}
}
