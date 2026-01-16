package services

import (
	"financialcontrol/internal/constants"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store"
	u "financialcontrol/internal/utils"
	ca "financialcontrol/internal/v1/categories/models"
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionsService struct {
	categoriesRepository   ca.CategoriesRepository
	creditcardsRepository  cr.CreditCardsRepository
	transactionsRepository tm.TransactionsRepository
}

func NewTransactionsService(
	categoriesRepository ca.CategoriesRepository,
	creditcardsRepository cr.CreditCardsRepository,
	transactionsRepository tm.TransactionsRepository,
) tm.TransactionsService {
	return TransactionsService{
		categoriesRepository:   categoriesRepository,
		creditcardsRepository:  creditcardsRepository,
		transactionsRepository: transactionsRepository,
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
	if request.CreditcardID != nil {
		creditcard, status, errs := t.readCreditcard(ctx, userID, *request.CreditcardID)

		if len(errs) > 0 {
			return tm.TransactionRelations{}, status, errs
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
	// TODO: validar se o cartão de crédito tem limite suficiente para a transação
	// TODO: remover strings hardcoded

	return tm.TransactionRelations{
		UserID:             userID,
		Request:            request,
		CategoryResponse:   categoryResponse,
		CreditcardResponse: creditcardResponse,
	}, http.StatusOK, nil
}

func (t TransactionsService) readCategory(ctx *gin.Context, userID uuid.UUID, id uuid.UUID) (cm.Category, int, []e.ApiError) {
	categoryNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CategoryNotFound}}

	category, errs := t.categoriesRepository.ReadByID(ctx, id)

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

	creditcard, errs := t.creditcardsRepository.ReadByID(ctx, id)

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

	transaction, errs := t.transactionsRepository.ReadById(ctx, transactionId)

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
