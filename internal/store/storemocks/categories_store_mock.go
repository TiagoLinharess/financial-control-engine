package storemocks

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type CategoriesStoreMock struct {
	Error                 error
	CategoryResult        pgstore.Category
	CategoriesResult      []pgstore.Category
	CategoriesCount       int64
	HasTransactionsResult bool
}

func NewCategoriesStoreMock() CategoriesStoreMock {
	return CategoriesStoreMock{
		Error:                 nil,
		CategoryResult:        pgstore.Category{},
		CategoriesResult:      []pgstore.Category{},
		CategoriesCount:       0,
		HasTransactionsResult: false,
	}
}

func (c CategoriesStoreMock) CreateCategory(ctx context.Context, arg pgstore.CreateCategoryParams) (pgstore.Category, error) {
	return c.CategoryResult, c.Error
}

func (c CategoriesStoreMock) GetCategoriesByUserID(ctx context.Context, userID uuid.UUID) ([]pgstore.Category, error) {
	return c.CategoriesResult, c.Error
}

func (c CategoriesStoreMock) CountCategoriesByUserID(ctx context.Context, userID uuid.UUID) (int64, error) {
	return c.CategoriesCount, c.Error
}

func (c CategoriesStoreMock) DeleteCategoryByID(ctx context.Context, id uuid.UUID) error {
	return c.Error
}

func (c CategoriesStoreMock) GetCategoryByID(ctx context.Context, id uuid.UUID) (pgstore.Category, error) {
	return c.CategoryResult, c.Error
}

func (c CategoriesStoreMock) UpdateCategory(ctx context.Context, arg pgstore.UpdateCategoryParams) (pgstore.Category, error) {
	return c.CategoryResult, c.Error
}

func (c CategoriesStoreMock) HasTransactionsByCategory(ctx context.Context, categoryID uuid.UUID) (bool, error) {
	if c.Error != nil {
		return false, c.Error
	}
	return c.HasTransactionsResult, nil
}
