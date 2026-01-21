package models

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type MonthlyTransactionsRepository interface {
	Create(request CreateMonthlyTransaction) (MonthlyTransaction, []e.ApiError)
	Read(m.PaginatedParams) ([]MonthlyTransaction, int64, []e.ApiError)
	ReadById(id uuid.UUID) (MonthlyTransaction, int, []e.ApiError)
	Update(MonthlyTransaction) (MonthlyTransaction, int, []e.ApiError)
	Delete(id uuid.UUID) (int, []e.ApiError)
}
