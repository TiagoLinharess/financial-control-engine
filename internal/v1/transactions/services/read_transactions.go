package services

import (
	"financialcontrol/internal/constants"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (t TransactionsService) Read(ctx *gin.Context) (m.PaginatedResponse[tm.TransactionResponse], int, []e.ApiError) {
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusUnauthorized, errs
	}

	limitString := ctx.DefaultQuery(constants.LimitText, constants.LimitDefaultString)
	limit, err := u.StringToInt64(limitString)

	if limit > constants.LimitDefault || err != nil {
		limit = constants.LimitDefault
	}

	pageString := ctx.DefaultQuery(constants.PageText, constants.PageDefaultString)
	page, err := u.StringToInt64(pageString)

	if err != nil {
		return m.PaginatedResponse[tm.TransactionResponse]{}, http.StatusBadRequest, []e.ApiError{e.CustomError{Message: constants.InvalidPageParam}}
	}

	if page == 0 {
		page = 1
	}

	offset := limit * (page - 1)

	startDateString := ctx.DefaultQuery(constants.StartDateText, constants.EmptyString)
	endDateString := ctx.DefaultQuery(constants.EndDateText, constants.EmptyString)

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
		errs = append(errs, e.CustomError{Message: constants.InvalidStartDate})
	}

	endDate, err := time.Parse(time.DateOnly, endDateString)

	if err != nil {
		errs = append(errs, e.CustomError{Message: constants.InvalidEndDate})
	}

	return startDate, endDate, errs
}
