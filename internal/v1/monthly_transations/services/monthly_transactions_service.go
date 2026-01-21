package services

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	mtm "financialcontrol/internal/v1/monthly_transations/models"

	"github.com/gin-gonic/gin"
)

type MonthlyTransactionsService struct {
	repository mtm.MonthlyTransactionsRepository
}

func NewMonthlyTransactionsService(repository mtm.MonthlyTransactionsRepository) *MonthlyTransactionsService {
	return &MonthlyTransactionsService{
		repository: repository,
	}
}

// Delete implements [models.MonthlyTransactionsService].
func (m *MonthlyTransactionsService) Delete(ctx *gin.Context) (int, []e.ApiError) {
	panic("unimplemented")
}

// Read implements [models.MonthlyTransactionsService].
func (m *MonthlyTransactionsService) Read(ctx *gin.Context) (m.PaginatedResponse[mtm.MonthlyTransactionResponse], int, []e.ApiError) {
	panic("unimplemented")
}

// ReadById implements [models.MonthlyTransactionsService].
func (m *MonthlyTransactionsService) ReadById(ctx *gin.Context) (mtm.MonthlyTransactionResponse, int, []e.ApiError) {
	panic("unimplemented")
}

// Update implements [models.MonthlyTransactionsService].
func (m *MonthlyTransactionsService) Update(ctx *gin.Context) (mtm.MonthlyTransactionResponse, int, []e.ApiError) {
	panic("unimplemented")
}
