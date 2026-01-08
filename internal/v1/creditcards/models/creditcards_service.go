package models

import (
	"financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
)

type CreditCardsService interface {
	Create(ctx *gin.Context) (CreditCardResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) ([]CreditCardResponse, int, []errors.ApiError)
	ReadAt(ctx *gin.Context) (CreditCardResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (CreditCardResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}
