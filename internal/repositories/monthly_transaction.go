package repositories

import (
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

type MonthlyTransaction interface {
	CreateMonthlyTransaction(request models.CreateMonthlyTransaction) (models.MonthlyTransaction, error)
	ReadMonthlyTransactionsByUserIDPaginated(params commonsmodels.PaginatedParams) ([]models.MonthlyTransaction, int64, error)
	ReadMonthlyTransactionByID(id uuid.UUID) (models.MonthlyTransaction, int, error)
	UpdateMonthlyTransaction(model models.MonthlyTransaction) (models.MonthlyTransaction, int, error)
	DeleteMonthlyTransaction(id uuid.UUID) (int, error)
}

func (r Repository) CreateMonthlyTransaction(request models.CreateMonthlyTransaction) (models.MonthlyTransaction, error) {
	panic("unimplemented")
}

func (r Repository) DeleteMonthlyTransaction(id uuid.UUID) (int, error) {
	panic("unimplemented")
}

func (r Repository) ReadMonthlyTransactionsByUserIDPaginated(params commonsmodels.PaginatedParams) ([]models.MonthlyTransaction, int64, error) {
	panic("unimplemented")
}

func (r Repository) ReadMonthlyTransactionByID(id uuid.UUID) (models.MonthlyTransaction, int, error) {
	panic("unimplemented")
}

func (r Repository) UpdateMonthlyTransaction(model models.MonthlyTransaction) (models.MonthlyTransaction, int, error) {
	panic("unimplemented")
}
