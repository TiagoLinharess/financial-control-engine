package models

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"net/http"
)

type CategoriesService interface {
	CreateCategory(w http.ResponseWriter, r *http.Request) (CategoryResponse, int, []errors.ApiError)
	ReadCategoriesByUser(w http.ResponseWriter, r *http.Request) (models.ResponseList[CategoryResponse], int, []errors.ApiError)
	ReadCategoryByID(w http.ResponseWriter, r *http.Request) (CategoryResponse, int, []errors.ApiError)
}
