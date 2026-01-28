package repositories

import (
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

type MonthlyTransaction interface {
	CreateMonthlyTransaction(request models.CreateMonthlyTransaction) (models.MonthlyTransaction, []errors.ApiError)
	ReadMonthlyTransactionsByUserIDPaginated(params models.PaginatedParams) ([]models.MonthlyTransaction, int64, []errors.ApiError)
	ReadMonthlyTransactionByID(id uuid.UUID) (models.MonthlyTransaction, int, []errors.ApiError)
	UpdateMonthlyTransaction(model models.MonthlyTransaction) (models.MonthlyTransaction, int, []errors.ApiError)
	DeleteMonthlyTransaction(id uuid.UUID) (int, []errors.ApiError)
}

func (r Repository) CreateMonthlyTransaction(request models.CreateMonthlyTransaction) (models.MonthlyTransaction, []errors.ApiError) {
	panic("unimplemented")
}

func (r Repository) DeleteMonthlyTransaction(id uuid.UUID) (int, []errors.ApiError) {
	panic("unimplemented")
}

func (r Repository) ReadMonthlyTransactionsByUserIDPaginated(params models.PaginatedParams) ([]models.MonthlyTransaction, int64, []errors.ApiError) {
	panic("unimplemented")
}

func (r Repository) ReadMonthlyTransactionByID(id uuid.UUID) (models.MonthlyTransaction, int, []errors.ApiError) {
	panic("unimplemented")
}

func (r Repository) UpdateMonthlyTransaction(model models.MonthlyTransaction) (models.MonthlyTransaction, int, []errors.ApiError) {
	panic("unimplemented")
}
