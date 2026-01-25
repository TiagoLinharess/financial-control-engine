package categories

import (
	"financialcontrol/internal/constants"
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"
	"financialcontrol/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (models.ResponseList[CategoryResponse], int, []errors.ApiError)
	ReadByID(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return service{repository: repository}
}

func (c service) Create(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.GetCategoryCountByUser(ctx, userID)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return CategoryResponse{}, http.StatusForbidden, []errors.ApiError{errors.LimitError{Message: errors.CategoriesLimit}}
	}

	request, errs := utils.DecodeValidJson[CategoryRequest](ctx)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusBadRequest, errs
	}

	data := request.ToCreateModel(userID)

	category, errs := c.repository.CreateCategory(ctx, data)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return category.ToResponse(), http.StatusCreated, nil
}

func (c service) Read(ctx *gin.Context) (models.ResponseList[CategoryResponse], int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.ResponseList[CategoryResponse]{}, http.StatusUnauthorized, errs
	}

	categories, errs := c.repository.ReadCategories(ctx, userID)

	if len(errs) > 0 {
		return models.ResponseList[CategoryResponse]{}, http.StatusInternalServerError, errs
	}

	response := make([]CategoryResponse, 0, len(categories))
	for _, category := range categories {
		response = append(response, category.ToResponse())
	}

	return models.ResponseList[CategoryResponse]{Items: response, Total: len(response)}, http.StatusOK, nil
}

func (s service) ReadByID(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return CategoryResponse{}, statusCode, errs
	}

	return category.ToResponse(), http.StatusOK, nil
}

func (s service) Update(ctx *gin.Context) (CategoryResponse, int, []errors.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return CategoryResponse{}, statusCode, errs
	}

	request, errs := utils.DecodeValidJson[CategoryRequest](ctx)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusBadRequest, errs
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, errs := s.repository.UpdateCategory(ctx, category)

	if len(errs) > 0 {
		return CategoryResponse{}, http.StatusInternalServerError, errs
	}

	return categoryEdited.ToResponse(), http.StatusOK, nil
}

func (s service) Delete(ctx *gin.Context) (int, []errors.ApiError) {
	category, statusCode, errs := s.read(ctx)

	if len(errs) > 0 {
		return statusCode, errs
	}

	hasTransactions, errs := s.repository.HasTransactionsByCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	if hasTransactions {
		return http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.CategoryCannotBeDeletedMsg}}
	}

	errs = s.repository.DeleteCategory(ctx, category.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}

func (s service) read(ctx *gin.Context) (Category, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return Category{}, http.StatusUnauthorized, errs
	}

	categoryIDString := ctx.Param(constants.ID)

	categoryID, err := uuid.Parse(categoryIDString)

	if err != nil {
		return Category{}, http.StatusBadRequest, errs
	}

	category, errs := s.repository.ReadCategoryByID(ctx, categoryID)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return Category{}, http.StatusNotFound, categoryNotFoundErr
		}
		return Category{}, http.StatusInternalServerError, errs
	}

	if category.UserID != userID {
		return Category{}, http.StatusNotFound, categoryNotFoundErr
	}

	return category, http.StatusOK, nil
}
