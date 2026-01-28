package services

import (
	"context"
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/constants"
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"
	"financialcontrol/internal/modelsdto"
	"financialcontrol/internal/repositories"
	"net/http"

	"github.com/google/uuid"
)

type Category interface {
	Create(ctx context.Context, userID uuid.UUID, request dtos.CategoryRequest) (dtos.CategoryResponse, errors.ApiError)
	Read(ctx context.Context, userID uuid.UUID) (commonsmodels.ResponseList[dtos.CategoryResponse], errors.ApiError)
	ReadByID(ctx context.Context, userID uuid.UUID, id uuid.UUID) (dtos.CategoryResponse, errors.ApiError)
	Update(ctx context.Context, userID uuid.UUID, id uuid.UUID, request dtos.CategoryRequest) (dtos.CategoryResponse, errors.ApiError)
	Delete(ctx context.Context, userID uuid.UUID, id uuid.UUID) errors.ApiError
}

type category struct {
	repository repositories.Category
}

func NewCategoriesService(repository repositories.Category) Category {
	return category{
		repository: repository,
	}
}

func (c category) Create(ctx context.Context, userID uuid.UUID, request dtos.CategoryRequest) (dtos.CategoryResponse, errors.ApiError) {
	count, err := c.repository.GetCategoryCountByUser(ctx, userID)

	if err != nil {
		return dtos.CategoryResponse{}, errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	if count >= 10 {
		return dtos.CategoryResponse{}, errors.NewApiError(
			http.StatusForbidden,
			errors.BadRequestError(constants.LimitReachedMsg),
		)
	}

	data := modelsdto.CreateCategoryFromRequest(request, userID)

	category, err := c.repository.CreateCategory(ctx, data)

	if err != nil {
		return dtos.CategoryResponse{}, errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	return modelsdto.CategoryResponseFromModel(category), nil
}

func (c category) Read(ctx context.Context, userID uuid.UUID) (commonsmodels.ResponseList[dtos.CategoryResponse], errors.ApiError) {
	categories, err := c.repository.ReadCategories(ctx, userID)

	if err != nil {
		return commonsmodels.ResponseList[dtos.CategoryResponse]{}, errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	response := make([]dtos.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		response = append(response, modelsdto.CategoryResponseFromModel(category))
	}

	return commonsmodels.ResponseList[dtos.CategoryResponse]{Items: response, Total: len(response)}, nil
}

func (c category) ReadByID(ctx context.Context, userID uuid.UUID, id uuid.UUID) (dtos.CategoryResponse, errors.ApiError) {
	category, apiErr := c.read(ctx, userID, id)

	if apiErr != nil {
		return dtos.CategoryResponse{}, apiErr
	}

	return modelsdto.CategoryResponseFromModel(category), nil
}

func (c category) Update(ctx context.Context, userID uuid.UUID, id uuid.UUID, request dtos.CategoryRequest) (dtos.CategoryResponse, errors.ApiError) {
	category, apiErr := c.read(ctx, userID, id)

	if apiErr != nil {
		return dtos.CategoryResponse{}, apiErr
	}

	category.Icon = request.Icon
	category.Name = request.Name
	category.TransactionType = *request.TransactionType

	categoryEdited, err := c.repository.UpdateCategory(ctx, category)

	if err != nil {
		return dtos.CategoryResponse{}, errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	return modelsdto.CategoryResponseFromModel(categoryEdited), nil
}

func (c category) Delete(ctx context.Context, userID uuid.UUID, id uuid.UUID) errors.ApiError {
	category, apiErr := c.read(ctx, userID, id)
	if apiErr != nil {
		return apiErr
	}

	hasTransactions, err := c.repository.HasTransactionsByCategory(ctx, category.ID)

	if err != nil {
		return errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	if hasTransactions {
		return errors.NewApiError(
			http.StatusBadRequest,
			errors.BadRequestError(constants.CannotBeDeletedMsg),
		)
	}

	err = c.repository.DeleteCategory(ctx, category.ID)

	if err != nil {
		return errors.NewApiError(
			http.StatusInternalServerError,
			errors.InternalServerError(err.Error()),
		)
	}

	return nil
}

func (c category) read(ctx context.Context, userID uuid.UUID, id uuid.UUID) (models.Category, errors.ApiError) {
	category, err := c.repository.ReadCategoryByID(ctx, id)

	if err != nil {
		if err.Error() == constants.StoreErrorNoRowsMsg {
			return models.Category{}, errors.NewApiError(http.StatusNotFound, errors.NotFoundError())
		}
		return models.Category{}, errors.NewApiError(http.StatusInternalServerError, errors.InternalServerError(err.Error()))
	}

	if category.UserID != userID {
		return models.Category{}, errors.NewApiError(http.StatusNotFound, errors.NotFoundError())
	}

	return category, nil
}
