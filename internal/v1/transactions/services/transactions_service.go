package services

import (
	"financialcontrol/internal/constants"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store/models"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionsService struct {
	repository st.TransactionsRepository
}

func NewTransactionsService(
	repository st.TransactionsRepository,
) tm.TransactionsService {
	return TransactionsService{
		repository: repository,
	}
}

func (t TransactionsService) getRelations(ctx *gin.Context) (tm.TransactionRelations, int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return tm.TransactionRelations{}, http.StatusUnauthorized, errs
	}

	request, errs := u.DecodeValidJson[tm.TransactionRequest](ctx)

	if len(errs) > 0 {
		return tm.TransactionRelations{}, http.StatusBadRequest, errs
	}

	category, status, errs := t.readCategory(ctx, userID, request.CategoryID)
	categoryResponse := category.ToShortResponse()

	if len(errs) > 0 {
		return tm.TransactionRelations{}, status, errs
	}

	var creditcardResponse *cr.ShortCreditCardResponse
	var creditcard *cr.CreditCard
	if request.CreditcardID != nil {
		creditcard, status, errs = t.readCreditcard(ctx, userID, *request.CreditcardID)

		if len(errs) > 0 {
			return tm.TransactionRelations{}, status, errs
		}

		totalAmountModel := tm.TransactionsCreditCardTotal{
			Date:         request.Date,
			UserID:       userID,
			CreditcardID: creditcard.ID,
		}

		totalAmount, err := t.repository.GetCreditcardTotalAmount(ctx, totalAmountModel)

		if err != nil {
			return tm.TransactionRelations{}, http.StatusInternalServerError, []e.ApiError{e.CustomError{Message: constants.InternalServerErrorMsg}}
		}

		if totalAmount+request.Value > creditcard.Limit {
			return tm.TransactionRelations{}, http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.TransactionCreditcardLimitExceededMsg}}
		}

		resp := creditcard.ToShortResponse()
		creditcardResponse = &resp
	}

	if request.CreditcardID == nil && category.TransactionType == m.Credit {
		return tm.TransactionRelations{}, http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.TransactionCreditWithoutCreditcardMsg}}
	}

	if request.CreditcardID != nil && category.TransactionType != m.Credit {
		return tm.TransactionRelations{}, http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.TransactionDebitOrIncomeWithCreditcardMsg}}
	}

	// TODO: validar assim como no cartão de crédito, as despesas mensais, anuais e parceladas

	return tm.TransactionRelations{
		UserID:             userID,
		Request:            request,
		CategoryResponse:   categoryResponse,
		CreditcardResponse: creditcardResponse,
	}, http.StatusOK, nil
}

func (t TransactionsService) readCategory(ctx *gin.Context, userID uuid.UUID, id uuid.UUID) (cm.Category, int, []e.ApiError) {
	categoryNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CategoryNotFound}}

	category, errs := t.repository.ReadCategoryByID(ctx, id)

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

func (t TransactionsService) readCreditcard(ctx *gin.Context, userID uuid.UUID, id uuid.UUID) (*cr.CreditCard, int, []e.ApiError) {
	creditcardNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CreditcardNotFound}}

	creditcard, errs := t.repository.ReadCreditCardByID(ctx, id)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == string(st.ErrNoRows)
		})
		if isNotFoundErr {
			return &cr.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
		}
		return &cr.CreditCard{}, http.StatusInternalServerError, errs
	}

	if creditcard.UserID != userID {
		return &cr.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
	}

	return &creditcard, http.StatusOK, nil
}

func (t TransactionsService) read(ctx *gin.Context) (tm.Transaction, int, []e.ApiError) {
	transactionNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.TransactionNotFound}}
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return tm.Transaction{}, http.StatusUnauthorized, errs
	}

	transactionIdString := ctx.Param(constants.ID)

	transactionId, err := uuid.Parse(transactionIdString)

	if err != nil {
		return tm.Transaction{}, http.StatusBadRequest, errs
	}

	transaction, errs := t.repository.ReadTransactionById(ctx, transactionId)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == string(st.ErrNoRows)
		})
		if isNotFoundErr {
			return tm.Transaction{}, http.StatusNotFound, transactionNotFoundErr
		}
		return tm.Transaction{}, http.StatusInternalServerError, errs
	}

	if transaction.UserID != userID {
		return tm.Transaction{}, http.StatusNotFound, transactionNotFoundErr
	}

	return transaction, http.StatusOK, nil
}
