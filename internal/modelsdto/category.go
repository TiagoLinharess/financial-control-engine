package modelsdto

import (
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func CreateCategoryFromRequest(dto dtos.CategoryRequest, userID uuid.UUID) models.CreateCategory {
	return models.CreateCategory{
		UserID:          userID,
		TransactionType: *dto.TransactionType,
		Name:            dto.Name,
		Icon:            dto.Icon,
	}
}

func CategoryResponseFromModel(model models.Category) dtos.CategoryResponse {
	return dtos.CategoryResponse{
		ID:              model.ID,
		TransactionType: model.TransactionType,
		Name:            model.Name,
		Icon:            model.Icon,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
	}
}

func ShortCategoryResponseFromModel(model models.Category) dtos.ShortCategoryResponse {
	return dtos.ShortCategoryResponse{
		ID:              model.ID,
		TransactionType: model.TransactionType,
		Name:            model.Name,
		Icon:            model.Icon,
	}
}

func ShortCategoryResponseFromShortModel(model models.ShortCategory) dtos.ShortCategoryResponse {
	return dtos.ShortCategoryResponse{
		ID:              model.ID,
		TransactionType: model.TransactionType,
		Name:            model.Name,
		Icon:            model.Icon,
	}
}
