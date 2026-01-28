package handlers

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/services"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type MonthlyTransaction struct {
	service services.MonthlyTransaction
}

func NewMonthlyTransactionsHandler(service services.MonthlyTransaction) *MonthlyTransaction {
	return &MonthlyTransaction{
		service: service,
	}
}

func (mc *MonthlyTransaction) Create(ctx *gin.Context) {
	response, status, errs := mc.service.Create(ctx)
	utils.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransaction) Read(ctx *gin.Context) {
	response, status, errs := mc.service.Read(ctx)
	utils.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransaction) ReadById(ctx *gin.Context) {
	response, status, errs := mc.service.ReadById(ctx)
	utils.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransaction) Update(ctx *gin.Context) {
	response, status, errs := mc.service.Update(ctx)
	utils.SendResponse(ctx, response, status, errs)
}

func (mc *MonthlyTransaction) Delete(ctx *gin.Context) {
	status, errs := mc.service.Delete(ctx)
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, errs)
}
