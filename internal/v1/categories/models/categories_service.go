package models

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
)

type CategoriesService interface {
	Create(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (models.ResponseList[CategoryResponse], int, []errors.ApiError)
	ReadByID(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}
