package models

import (
	"financialcontrol/internal/models/errors"
	"net/http"
)

type CategoriesService interface {
	CreateCategory(w http.ResponseWriter, r *http.Request) (CategoryResponse, int, []errors.ApiError)
}
