package models

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
)

type TransactionsService interface {
	Create(ctx *gin.Context) (TransactionResponse, int, []e.ApiError)
	Read(ctx *gin.Context) (m.PaginatedResponse[TransactionResponse], int, []e.ApiError)
	ReadById(ctx *gin.Context) (TransactionResponse, int, []e.ApiError)
	Update(ctx *gin.Context) (TransactionResponse, int, []e.ApiError)
	Delete(ctx *gin.Context) (int, []e.ApiError)
	Pay(ctx *gin.Context) (int, []e.ApiError)
}
