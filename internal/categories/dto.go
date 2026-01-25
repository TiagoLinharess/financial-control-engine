package categories

import "github.com/google/uuid"

func (c CategoryRequest) ToCreateModel(userID uuid.UUID) CreateCategory {
	return CreateCategory{
		UserID:          userID,
		TransactionType: *c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}

func (c Category) ToResponse() CategoryResponse {
	return CategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}

func (c Category) ToShortResponse() ShortCategoryResponse {
	return ShortCategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}

func (c ShortCategory) ToShortResponse() ShortCategoryResponse {
	return ShortCategoryResponse{
		ID:              c.ID,
		TransactionType: c.TransactionType,
		Name:            c.Name,
		Icon:            c.Icon,
	}
}
