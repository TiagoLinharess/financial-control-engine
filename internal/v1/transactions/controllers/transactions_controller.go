package controllers

import (
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
