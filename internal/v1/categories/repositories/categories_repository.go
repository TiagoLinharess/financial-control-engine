package repositories

import (
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/v1/categories/models"
	"time"

	"github.com/google/uuid"
)

type CategoriesRepository struct{}

func NewCategoriesRepository() models.CategoriesRepository {
	return CategoriesRepository{}
}

func (c CategoriesRepository) CreateCategory(data models.CreateCategory) (models.Category, []errors.ApiError) {
	return models.Category{
		ID:              uuid.New(),
		UserID:          data.UserID,
		TransactionType: data.TransactionType,
		Name:            data.Name,
		Icon:            data.Icon,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}
