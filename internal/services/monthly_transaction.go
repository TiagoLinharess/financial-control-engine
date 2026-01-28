package services

import (
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/modelsdto"
	"financialcontrol/internal/repositories"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MonthlyTransaction interface {
	Create(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (commonsmodels.PaginatedResponse[dtos.MonthlyTransactionResponse], int, []errors.ApiError)
	ReadById(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}

type monthlyTransaction struct {
	repository repositories.MonthlyTransaction
}

func NewMonthlyTransactionService(repository repositories.MonthlyTransaction) MonthlyTransaction {
	return &monthlyTransaction{
		repository: repository,
	}
}

// Create implements [MonthlyTransaction].
func (m *monthlyTransaction) Create(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError) {
	userId, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return dtos.MonthlyTransactionResponse{}, 0, errs
	}

	request, errs := utils.DecodeValidJson[dtos.MonthlyTransactionRequest](ctx)

	if len(errs) > 0 {
		return dtos.MonthlyTransactionResponse{}, 0, errs
	}

	createModel := modelsdto.CreateMonthlyTransactionFromRequest(request, userId)

	model, errs := m.repository.CreateMonthlyTransaction(createModel)

	if len(errs) > 0 {
		return dtos.MonthlyTransactionResponse{}, 0, errs
	}

	response := modelsdto.MonthlyTransactionResponseFromModel(model)

	return response, http.StatusCreated, nil
}

// Delete implements [MonthlyTransaction].
func (m *monthlyTransaction) Delete(ctx *gin.Context) (int, []errors.ApiError) {
	panic("unimplemented")
}

// Read implements [MonthlyTransaction].
func (m *monthlyTransaction) Read(ctx *gin.Context) (commonsmodels.PaginatedResponse[dtos.MonthlyTransactionResponse], int, []errors.ApiError) {
	panic("unimplemented")
}

// ReadById implements [MonthlyTransaction].
func (m *monthlyTransaction) ReadById(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError) {
	panic("unimplemented")
}

// Update implements [MonthlyTransaction].
func (m *monthlyTransaction) Update(ctx *gin.Context) (dtos.MonthlyTransactionResponse, int, []errors.ApiError) {
	panic("unimplemented")
}
