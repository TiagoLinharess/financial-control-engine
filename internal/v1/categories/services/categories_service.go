package services

import (
	"financialcontrol/internal/constants"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoriesService struct {
	repository cm.CategoriesRepository
}

func NewCategoriesService(repository cm.CategoriesRepository) cm.CategoriesService {
	return CategoriesService{repository: repository}
}

func (s CategoriesService) read(ctx *gin.Context) (cm.Category, int, []e.ApiError) {
	categoryNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CategoryNotFound}}
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return cm.Category{}, http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param(constants.ID)

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return cm.Category{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadByID(ctx, categoryID)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == string(st.ErrNoRows)
		})
		if isNotFoundErr {
			return cm.Category{}, http.StatusNotFound, categoryNotFoundErr
		}
		return cm.Category{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return cm.Category{}, http.StatusNotFound, categoryNotFoundErr
	}

	return category, http.StatusOK, nil
}
