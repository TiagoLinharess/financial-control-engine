package controllers

import (
	m "financialcontrol/internal/models"
	u "financialcontrol/internal/utils"
	mtm "financialcontrol/internal/v1/monthly_transations/models"

	"github.com/gin-gonic/gin"
)

type MonthlyTransactionsController struct {
	service mtm.MonthlyTransactionsService
}

func NewMonthlyTransactionsController(service mtm.MonthlyTransactionsService) *MonthlyTransactionsController {
	return &MonthlyTransactionsController{
		service: service,
	}
}

func (mc *MonthlyTransactionsController) Create(ctx *gin.Context) {
	response, status, errs := mc.service.Create(ctx)
	u.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransactionsController) Read(ctx *gin.Context) {
	response, status, errs := mc.service.Read(ctx)
	u.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransactionsController) ReadById(ctx *gin.Context) {
	response, status, errs := mc.service.ReadById(ctx)
	u.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransactionsController) Update(ctx *gin.Context) {
	response, status, errs := mc.service.Update(ctx)
	u.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransactionsController) Delete(ctx *gin.Context) {
	status, errs := mc.service.Delete(ctx)
	u.SendResponse(ctx, m.NewResponseSuccess(), status, errs)
}
