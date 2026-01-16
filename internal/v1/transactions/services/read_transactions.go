package services

import (
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"
	"strconv"
	"time"

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

	startDateString := ctx.DefaultQuery("start_date", "")
	endDateString := ctx.DefaultQuery("end_date", "")

	var responses []tm.Transaction
	var count int64

	if !u.IsBlank(startDateString) || !u.IsBlank(endDateString) {
		startDate, endDate, errs := t.readDatesFrom(startDateString, endDateString)

		if len(errs) > 0 {
			return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusBadRequest, errs
		}

		paginatedParams := m.PaginatedParamsWithDateRange{
			UserID:    userID,
			Limit:     int32(limit),
			Offset:    int32(offset),
			StartDate: startDate,
			EndDate:   endDate,
		}

		responses, count, errs = t.transactionsRepository.ReadInToDates(ctx, paginatedParams)
	} else {
		paginatedParams := m.PaginatedParams{
			UserID: userID,
			Limit:  int32(limit),
			Offset: int32(offset),
		}

		responses, count, errs = t.transactionsRepository.Read(ctx, paginatedParams)
	}

	if len(errs) > 0 {
		return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusInternalServerError, errs
	}

	transactionsResponse := make([]tm.TransactionResponse, 0, len(responses))

	for _, transaction := range responses {
		transactionsResponse = append(transactionsResponse, transaction.ToResponse())
	}

	return m.PaginatedResponse[tm.TransactionResponse]{
		Items:     transactionsResponse,
		PageCount: (count / limit) + 1,
		Page:      page,
	}, http.StatusOK, nil
}

func (t TransactionsService) readDatesFrom(startDateString string, endDateString string) (time.Time, time.Time, []e.ApiError) {
	errs := make([]e.ApiError, 0)
	startDate, err := time.Parse(time.DateOnly, startDateString)

	if err != nil {
		errs = append(errs, e.CustomError{Message: "Invalid start date"})
	}

	endDate, err := time.Parse(time.DateOnly, endDateString)

	if err != nil {
		errs = append(errs, e.CustomError{Message: "Invalid end date"})
	}

	return startDate, endDate, errs
}
