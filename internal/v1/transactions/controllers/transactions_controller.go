package controllers

import (
	m "financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	service tm.TransactionsService
}

func NewTransactionsController(service tm.TransactionsService) *TransactionsController {
	return &TransactionsController{service: service}
}

func (c *TransactionsController) Create(ctx *gin.Context) {
	data, status, err := c.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *TransactionsController) Read(ctx *gin.Context) {
	data, status, err := c.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *TransactionsController) ReadById(ctx *gin.Context) {
	data, status, err := c.service.ReadById(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *TransactionsController) Update(ctx *gin.Context) {
	data, status, err := c.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *TransactionsController) Delete(ctx *gin.Context) {
	status, err := c.service.Delete(ctx)
	utils.SendResponse(ctx, m.NewResponseSuccess(), status, err)
}

func (c *TransactionsController) Pay(ctx *gin.Context) {
	status, err := c.service.Pay(ctx)
	utils.SendResponse(ctx, m.NewResponseSuccess(), status, err)
}
