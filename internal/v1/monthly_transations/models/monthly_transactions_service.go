package models

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
)

type MonthlyTransactionsService interface {
	Create(ctx *gin.Context) (MonthlyTransactionResponse, int, []e.ApiError)
	Read(ctx *gin.Context) (m.PaginatedResponse[MonthlyTransactionResponse], int, []e.ApiError)
	ReadById(ctx *gin.Context) (MonthlyTransactionResponse, int, []e.ApiError)
	Update(ctx *gin.Context) (MonthlyTransactionResponse, int, []e.ApiError)
	Delete(ctx *gin.Context) (int, []e.ApiError)
}
