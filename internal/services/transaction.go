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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Transaction interface {
	Create(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError)
	Read(ctx *gin.Context) (models.PaginatedResponse[dtos.TransactionResponse], int, []errors.ApiError)
	ReadById(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError)
	Update(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError)
	Delete(ctx *gin.Context) (int, []errors.ApiError)
	Pay(ctx *gin.Context) (int, []errors.ApiError)
}

type transaction struct {
	repository repositories.Transaction
}

func NewTransactionsService(repository repositories.Transaction) Transaction {
	return &transaction{
		repository: repository,
	}
}

func (t transaction) Create(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError) {
	relations, status, errs := t.getRelations(ctx)

	if len(errs) > 0 {
		return dtos.TransactionResponse{}, status, errs
	}

	data := modelsdto.CreateTransactionFromTransactionRequest(relations.Request, relations.UserID)

	transaction, errs := t.repository.CreateTransaction(ctx, data)

	if len(errs) > 0 {
		return dtos.TransactionResponse{}, http.StatusInternalServerError, errs
	}

	response := modelsdto.TransactionResponseFromShortTransaction(transaction, relations.CategoryResponse, relations.CreditcardResponse)

	return response, http.StatusCreated, nil
}

func (t transaction) Read(ctx *gin.Context) (models.PaginatedResponse[dtos.TransactionResponse], int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.PaginatedResponse[dtos.TransactionResponse]{}, http.StatusUnauthorized, errs
	}

	limitString := ctx.DefaultQuery(constants.LimitText, constants.LimitDefaultString)
	limit, err := utils.StringToInt64(limitString)
	if limit > constants.LimitDefault || err != nil {
		limit = constants.LimitDefault
	}

	pageString := ctx.DefaultQuery(constants.PageText, constants.PageDefaultString)
	page, err := utils.StringToInt64(pageString)

	if err != nil {
		return models.PaginatedResponse[dtos.TransactionResponse]{}, http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.InvalidPageParam}}
	}

	if page == 0 {
		page = 1
	}

	offset := limit * (page - 1)

	startDateString := ctx.DefaultQuery(constants.StartDateText, constants.EmptyString)
	endDateString := ctx.DefaultQuery(constants.EndDateText, constants.EmptyString)

	var responses []models.Transaction
	var count int64

	if !utils.IsBlank(startDateString) || !utils.IsBlank(endDateString) {
		startDate, endDate, errs := t.readDatesFrom(startDateString, endDateString)

		if len(errs) > 0 {
			return models.PaginatedResponse[dtos.TransactionResponse]{}, http.StatusBadRequest, errs
		}

		paginatedParams := models.PaginatedParamsWithDateRange{
			UserID:    userID,
			Limit:     int32(limit),
			Offset:    int32(offset),
			StartDate: startDate,
			EndDate:   endDate,
		}

		responses, count, errs = t.repository.ReadTransactionsInToDates(ctx, paginatedParams)
	} else {
		paginatedParams := models.PaginatedParams{
			UserID: userID,
			Limit:  int32(limit),
			Offset: int32(offset),
		}

		responses, count, errs = t.repository.ReadTransactions(ctx, paginatedParams)
	}

	if len(errs) > 0 {
		return models.PaginatedResponse[dtos.TransactionResponse]{}, http.StatusInternalServerError, errs
	}

	transactionsResponse := make([]dtos.TransactionResponse, 0, len(responses))

	for _, transaction := range responses {
		transactionsResponse = append(transactionsResponse, modelsdto.TransactionResponseFromTransaction(transaction))
	}

	return models.PaginatedResponse[dtos.TransactionResponse]{
		Items:     transactionsResponse,
		PageCount: (count / limit) + 1,
		Page:      page,
	}, http.StatusOK, nil
}

func (t transaction) readDatesFrom(startDateString string, endDateString string) (time.Time, time.Time, []errors.ApiError) {
	errs := make([]errors.ApiError, 0)
	startDate, err := time.Parse(time.DateOnly, startDateString)

	if err != nil {
		errs = append(errs, errors.CustomError{Message: constants.InvalidStartDate})
	}

	endDate, err := time.Parse(time.DateOnly, endDateString)

	if err != nil {
		errs = append(errs, errors.CustomError{Message: constants.InvalidEndDate})
	}

	return startDate, endDate, errs
}

func (t transaction) ReadById(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError) {
	transaction, status, errs := t.read(ctx)
	return modelsdto.TransactionResponseFromTransaction(transaction), status, errs
}

func (t transaction) Update(ctx *gin.Context) (dtos.TransactionResponse, int, []errors.ApiError) {
	relations, status, errs := t.getRelations(ctx)

	if len(errs) > 0 {
		return dtos.TransactionResponse{}, status, errs
	}

	transaction, status, errs := t.read(ctx)

	if len(errs) > 0 {
		return dtos.TransactionResponse{}, status, errs
	}

	var creditcard *models.ShortCreditCard
	if relations.CreditcardResponse != nil {
		creditcardModel := models.ShortCreditCard(*relations.CreditcardResponse)
		creditcard = &creditcardModel
	}

	transaction.Name = relations.Request.Name
	transaction.Date = relations.Request.Date
	transaction.Paid = relations.Request.Paid
	transaction.Category = models.ShortCategory(relations.CategoryResponse)
	transaction.Creditcard = creditcard

	transactionUpdated, errs := t.repository.UpdateTransaction(ctx, transaction)

	if len(errs) > 0 {
		return dtos.TransactionResponse{}, http.StatusInternalServerError, errs
	}

	response := modelsdto.TransactionResponseFromShortTransaction(transactionUpdated, relations.CategoryResponse, relations.CreditcardResponse)

	return response, http.StatusOK, nil
}

func (t transaction) Delete(ctx *gin.Context) (int, []errors.ApiError) {
	transaction, status, errs := t.read(ctx)

	if len(errs) > 0 {
		return status, errs
	}

	errs = t.repository.DeleteTransaction(ctx, transaction.ID)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}

func (t transaction) Pay(ctx *gin.Context) (int, []errors.ApiError) {
	transaction, status, errs := t.read(ctx)

	if len(errs) > 0 {
		return status, errs
	}

	errs = t.repository.PayTransaction(ctx, transaction.ID, !transaction.Paid)

	if len(errs) > 0 {
		return http.StatusInternalServerError, errs
	}

	return http.StatusNoContent, nil
}

func (t transaction) getRelations(ctx *gin.Context) (dtos.TransactionRelations, int, []errors.ApiError) {
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return dtos.TransactionRelations{}, http.StatusUnauthorized, errs
	}

	request, errs := utils.DecodeValidJson[dtos.TransactionRequest](ctx)

	if len(errs) > 0 {
		return dtos.TransactionRelations{}, http.StatusBadRequest, errs
	}

	category, status, errs := t.readCategory(ctx, userID, request.CategoryID)
	categoryResponse := modelsdto.ShortCategoryResponseFromModel(category)

	if len(errs) > 0 {
		return dtos.TransactionRelations{}, status, errs
	}

	var creditcardResponse *dtos.ShortCreditCardResponse
	var creditcard *models.CreditCard
	if request.CreditcardID != nil {
		creditcard, status, errs = t.readCreditcard(ctx, userID, *request.CreditcardID)

		if len(errs) > 0 {
			return dtos.TransactionRelations{}, status, errs
		}

		totalAmountModel := models.TransactionsCreditCardTotal{
			Date:         request.Date,
			UserID:       userID,
			CreditcardID: creditcard.ID,
		}

		totalAmount, err := t.repository.GetCreditcardTotalAmount(ctx, totalAmountModel)

		if err != nil {
			return dtos.TransactionRelations{}, http.StatusInternalServerError, []errors.ApiError{errors.CustomError{Message: constants.InternalServerErrorMsg}}
		}

		if totalAmount+request.Value > creditcard.Limit {
			return dtos.TransactionRelations{}, http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.TransactionCreditcardLimitExceededMsg}}
		}

		resp := modelsdto.ShortCreditCardResponseFromCreditCard(*creditcard)
		creditcardResponse = &resp
	}

	if request.CreditcardID == nil && category.TransactionType == models.Credit {
		return dtos.TransactionRelations{}, http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.TransactionCreditWithoutCreditcardMsg}}
	}

	if request.CreditcardID != nil && category.TransactionType != models.Credit {
		return dtos.TransactionRelations{}, http.StatusBadRequest, []errors.ApiError{errors.CustomError{Message: constants.TransactionDebitOrIncomeWithCreditcardMsg}}
	}

	// TODO: validar assim como no cartão de crédito, as despesas mensais, anuais e parceladas

	return dtos.TransactionRelations{
		UserID:             userID,
		Request:            request,
		CategoryResponse:   categoryResponse,
		CreditcardResponse: creditcardResponse,
	}, http.StatusOK, nil
}

func (t transaction) readCategory(ctx *gin.Context, userID uuid.UUID, id uuid.UUID) (models.Category, int, []errors.ApiError) {
	categoryNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CategoryNotFound}}

	category, errs := t.repository.ReadCategoryByID(ctx, id)

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

func (t transaction) readCreditcard(ctx *gin.Context, userID uuid.UUID, id uuid.UUID) (*models.CreditCard, int, []errors.ApiError) {
	creditcardNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.CreditcardNotFound}}

	creditcard, errs := t.repository.ReadCreditCardByID(ctx, id)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return &models.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
		}
		return &models.CreditCard{}, http.StatusInternalServerError, errs
	}

	if creditcard.UserID != userID {
		return &models.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
	}

	return &creditcard, http.StatusOK, nil
}

func (t transaction) read(ctx *gin.Context) (models.Transaction, int, []errors.ApiError) {
	transactionNotFoundErr := []errors.ApiError{errors.NotFoundError{Message: errors.TransactionNotFound}}
	userID, errs := utils.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return models.Transaction{}, http.StatusUnauthorized, errs
	}

	transactionIdString := ctx.Param(constants.ID)

	transactionId, err := uuid.Parse(transactionIdString)

	if err != nil {
		return models.Transaction{}, http.StatusBadRequest, errs
	}

	transaction, errs := t.repository.ReadTransactionById(ctx, transactionId)

	if len(errs) > 0 {
		isNotFoundErr := utils.FindIf(errs, func(err errors.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return models.Transaction{}, http.StatusNotFound, transactionNotFoundErr
		}
		return models.Transaction{}, http.StatusInternalServerError, errs
	}

	if transaction.UserID != userID {
		return models.Transaction{}, http.StatusNotFound, transactionNotFoundErr
	}

	return transaction, http.StatusOK, nil
}
