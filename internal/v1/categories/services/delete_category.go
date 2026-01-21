package services

import (
	"financialcontrol/internal/constants"
	e "financialcontrol/internal/models/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s CategoriesService) Delete(ctx *gin.Context) (int, []e.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return statusCode, errs
	}

	hasTransactions, errs := s.repository.HasTransactionsByCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	if hasTransactions {
		return http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.CategoryCannotBeDeletedMsg}}
	}

	errs = s.repository.DeleteCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}
