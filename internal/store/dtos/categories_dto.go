package dtos

import (
	m "financialcontrol/internal/models"
	pgs "financialcontrol/internal/store/pgstore"
	cm "financialcontrol/internal/v1/categories/models"
)

func StoreCategoryModelToCategory(category pgs.Category) cm.Category {
	return cm.Category{
		ID:              category.ID,
		UserID:          category.UserID,
		TransactionType: m.TransactionType(category.TransactionType),
		Name:            category.Name,
		Icon:            category.Icon,
		CreatedAt:       category.CreatedAt.Time,
		UpdatedAt:       category.UpdatedAt.Time,
	}
}
