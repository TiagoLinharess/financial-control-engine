package repositories

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	s "financialcontrol/internal/store"
	mtm "financialcontrol/internal/v1/monthly_transations/models"

	"github.com/google/uuid"
)

type MonthlyTransactionsRepository struct {
	store s.MonthlyTransactionsStore
}

func NewMonthlyTransactionsRepository(store s.MonthlyTransactionsStore) *MonthlyTransactionsRepository {
	return &MonthlyTransactionsRepository{
		store: store,
	}
}

// Create implements [models.MonthlyTransactionsRepository].
func (m *MonthlyTransactionsRepository) Create(request mtm.CreateMonthlyTransaction) (mtm.MonthlyTransaction, []e.ApiError) {
	panic("unimplemented")
}

// Delete implements [models.MonthlyTransactionsRepository].
func (m *MonthlyTransactionsRepository) Delete(id uuid.UUID) (int, []e.ApiError) {
	panic("unimplemented")
}

// Read implements [models.MonthlyTransactionsRepository].
func (m *MonthlyTransactionsRepository) Read(m.PaginatedParams) ([]mtm.MonthlyTransaction, int64, []e.ApiError) {
	panic("unimplemented")
}

// ReadById implements [models.MonthlyTransactionsRepository].
func (m *MonthlyTransactionsRepository) ReadById(id uuid.UUID) (mtm.MonthlyTransaction, int, []e.ApiError) {
	panic("unimplemented")
}

// Update implements [models.MonthlyTransactionsRepository].
func (m *MonthlyTransactionsRepository) Update(mtm.MonthlyTransaction) (mtm.MonthlyTransaction, int, []e.ApiError) {
	panic("unimplemented")
}
