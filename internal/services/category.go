package services

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"
	"financialcontrol/internal/modelsdto"
	"financialcontrol/internal/repositories"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Category interface {
	Create(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (models.ResponseList[dtos.CategoryResponse], int, []errors.ApiError)
	ReadByID(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}

type category struct {
	repository repositories.Category
}

func NewCategoriesService(repository repositories.Category) Category {
	return category{
		repository: repository,
	}
}

func (c category) Create(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.GetCategoryCountByUser(ctx, userID)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return dtos.CategoryResponse{}, http.StatusForbidden, []errors.ApiError{errors.LimitError{Message: errors.CategoriesLimit}}
	}

	request, errs := utils.DecodeValidJson[dtos.CategoryRequest](ctx)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusBadRequest, errs
	}

	data := modelsdto.CreateCategoryFromRequest(request, userID)

	category, errs := c.repository.CreateCategory(ctx, data)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return modelsdto.CategoryResponseFromModel(category), http.StatusCreated, nil
}

func (c category) Read(ctx *gin.Context) (models.ResponseList[dtos.CategoryResponse], int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.ResponseList[dtos.CategoryResponse]{}, http.StatusUnauthorized, errs
	}

	categories, errs := c.repository.ReadCategories(ctx, userID)

	if len(errs) > 0 {
		return models.ResponseList[dtos.CategoryResponse]{}, http.StatusInternalServerError, errs
	}

	response := make([]dtos.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		response = append(response, modelsdto.CategoryResponseFromModel(category))
	}

	return models.ResponseList[dtos.CategoryResponse]{Items: response, Total: len(response)}, http.StatusOK, nil
}

func (c category) ReadByID(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError) {
	category, statusCode, errs := c.read(ctx)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, statusCode, errs
	}

	return modelsdto.CategoryResponseFromModel(category), http.StatusOK, nil
}

func (c category) Update(ctx *gin.Context) (dtos.CategoryResponse, int, []errors.ApiError) {
	category, statusCode, errs := c.read(ctx)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, statusCode, errs
	}

	request, errs := utils.DecodeValidJson[dtos.CategoryRequest](ctx)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusBadRequest, errs
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, errs := c.repository.UpdateCategory(ctx, category)

	if len(errs) > 0 {
		return dtos.CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return modelsdto.CategoryResponseFromModel(categoryEdited), http.StatusOK, nil
}

func (c category) Delete(ctx *gin.Context) (int, []errors.ApiError) {
	category, statusCode, errs := c.read(ctx)
	if len(errs) > 0 {
		return statusCode, errs
	}

	hasTransactions, errs := c.repository.HasTransactionsByCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	if hasTransactions {
		return http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.CategoryCannotBeDeletedMsg}}
	}

	errs = c.repository.DeleteCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}

func (c category) read(ctx *gin.Context) (models.Category, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.Category{}, http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param(constants.ID)

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return models.Category{}, http.StatusBadRequest, errs
	}

	category, errs := c.repository.ReadCategoryByID(ctx, categoryID)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return models.Category{}, http.StatusNotFound, categoryNotFoundErr
		}
		return models.Category{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return models.Category{}, http.StatusNotFound, categoryNotFoundErr
	}

	return category, http.StatusOK, nil
}
