package services

import (
	"financialcontrol/internal/models/errors"
	"net/http"
)

func (c CategoriesService) ReadCategoriesByUser(w http.ResponseWriter, r *http.Request) (int, []errors.ApiError) {

}
