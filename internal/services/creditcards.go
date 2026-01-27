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

type CreditCard interface {
	Create(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (models.ResponseList[dtos.CreditCardResponse], int, []errors.ApiError)
	ReadAt(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
}

type creditCard struct {
	repository repositories.CreditCard
}

func NewCreditCardService(repository repositories.CreditCard) CreditCard {
	return &creditCard{repository: repository}
}

func (c creditCard) Create(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return dtos.CreditCardResponse{}, http.StatusUnauthorized, errs
	}

	count, errs := c.repository.ReadCountByUser(ctx, userID)

	if len(errs) > 0 {
		return dtos.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	if count >= 10 {
		return dtos.CreditCardResponse{}, http.StatusForbidden, []errors.ApiError{errors.LimitError{Message: errors.CreditcardsLimit}}
	}

	request, errs := utils.DecodeValidJson[dtos.CreditCardRequest](ctx)

	if len(errs) > 0 {
		return dtos.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	model := modelsdto.CreditCardRequestToCreateModel(request, userID)

	creditCard, errs := c.repository.CreateCreditCard(ctx, model)

	if len(errs) > 0 {
		return dtos.CreditCardResponse{}, http.StatusInternalServerError, errs
	}

	return modelsdto.CreditCardToResponse(creditCard), http.StatusCreated, nil
}

func (c creditCard) Read(ctx *gin.Context) (models.ResponseList[dtos.CreditCardResponse], int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.ResponseList[dtos.CreditCardResponse]{}, http.StatusUnauthorized, errs
	}

	creditCards, errs := c.repository.ReadCreditCards(ctx, userID)

	if len(errs) > 0 {
		return models.ResponseList[dtos.CreditCardResponse]{}, http.StatusInternalServerError, errs
	}

	creditCardsResponse := make([]dtos.CreditCardResponse, 0, len(creditCards))
	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, modelsdto.CreditCardToResponse(creditCard))
	}

	return models.ResponseList[dtos.CreditCardResponse]{
		Items: creditCardsResponse,
		Total: len(creditCardsResponse),
	}, http.StatusOK, nil
}

func (c creditCard) ReadAt(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError) {
	creditcard, status, err := c.read(ctx)
	return modelsdto.CreditCardToResponse(creditcard), status, err
}

func (c creditCard) Update(ctx *gin.Context) (dtos.CreditCardResponse, int, []errors.ApiError) {
	creditcard, status, err := c.read(ctx)

	if len(err) > 0 {
		return dtos.CreditCardResponse{}, status, err
	}

	request, errs := utils.DecodeValidJson[dtos.CreditCardRequest](ctx)

	if len(errs) > 0 {
		return dtos.CreditCardResponse{}, http.StatusBadRequest, errs
	}

	creditcard.Name = request.Name
	creditcard.FirstFourNumbers = request.FirstFourNumbers
	creditcard.Limit = request.Limit
	creditcard.CloseDay = request.CloseDay
	creditcard.ExpireDay = request.ExpireDay
	creditcard.BackgroundColor = request.BackgroundColor
	creditcard.TextColor = request.TextColor

	creditcard, err = c.repository.UpdateCreditCard(ctx, creditcard)

	return modelsdto.CreditCardToResponse(creditcard), http.StatusOK, nil
}

func (c creditCard) Delete(ctx *gin.Context) (int, []errors.ApiError) {
	creditcard, status, err := c.read(ctx)

	if len(err) > 0 {
		return status, err
	}

	hasTransactions, err := c.repository.HasTransactionsByCreditCard(ctx, creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	if hasTransactions {
		return http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.CreditcardCannotBeDeletedMsg}}
	}

	err = c.repository.DeleteCreditCard(ctx, creditcard.ID)

	if len(err) > 0 {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

func (c creditCard) read(ctx *gin.Context) (models.CreditCard, int, []errors.ApiError) {
	creditcardNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CreditcardNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.CreditCard{}, http.StatusUnauthorized, errs
	}

	creditcardIdString := ctx.Param(constants.ID)

	creditcardId, err := uuid.Parse(creditcardIdString)

	if err != nil {
		return models.CreditCard{}, http.StatusBadRequest, errs
	}

	creditcard, errs := c.repository.ReadCreditCardByID(ctx, creditcardId)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return models.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
		}
		return models.CreditCard{}, http.StatusInternalServerError, errs
	}

	if creditcard.UserID != userID {
		return models.CreditCard{}, http.StatusForbidden, creditcardNotFoundErr
	}

	return creditcard, http.StatusOK, nil
}
