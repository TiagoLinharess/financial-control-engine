package services

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) Read(ctx *gin.Context) (m.PaginatedResponse[tm.TransactionResponse], int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusUnauthorized, errs
	}

	limitString := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitString, 10, 64)

	if limit > 10 || err != nil {
		limit = 10
	}

	pageString := ctx.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageString, 10, 64)

	if err != nil {
		return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusBadRequest, []e.ApiError{e.CustomError{Message: "Invalid page param"}}
	}

	if page == 0 {
		page = 1
	}

	offset := limit * (page - 1)

	paginatedParams := m.PaginatedParams{
		UserID: userID,
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	responses, count, errs := t.transactionsRepository.Read(ctx, paginatedParams)

	transactionsResponse := make([]tm.TransactionResponse, 0, len(responses))

	for _, transaction := range responses {
		transactionsResponse = append(transactionsResponse, transaction.ToResponse())
	}

	return m.PaginatedResponse[tm.TransactionResponse]{
		Items:     transactionsResponse,
		PageCount: count / limit,
		Page:      page,
	}, http.StatusOK, nil
}
