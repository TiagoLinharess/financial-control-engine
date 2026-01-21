package services

import (
	e "financialcontrol/internal/models/errors"
	u "financialcontrol/internal/utils"
	mtm "financialcontrol/internal/v1/monthly_transations/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MonthlyTransactionsService) Create(ctx *gin.Context) (mtm.MonthlyTransactionResponse, int, []e.ApiError) {
	userId, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return mtm.MonthlyTransactionResponse{}, 0, errs
	}

	request, errs := u.DecodeValidJson[mtm.MonthlyTransactionRequest](ctx)

	if len(errs) > 0 {
		return mtm.MonthlyTransactionResponse{}, 0, errs
	}

	createModel := request.ToCreateModel(userId)

	model, errs := m.repository.Create(createModel)

	if len(errs) > 0 {
		return mtm.MonthlyTransactionResponse{}, 0, errs
	}

	response := model.ToResponse()

	return response, http.StatusCreated, nil
}
