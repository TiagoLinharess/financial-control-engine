package models

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"

	"github.com/gin-gonic/gin"
)

type CategoriesService interface {
	CreateCategory(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	ReadCategoriesByUser(ctx *gin.Context) (models.ResponseList[CategoryResponse], int, []errors.ApiError)
	ReadCategoryByID(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	UpdateCategory(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	DeleteCategory(ctx *gin.Context) (int, []errors.ApiError)
}
