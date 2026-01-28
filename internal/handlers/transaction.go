package handlers

import (
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/services"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	service services.Transaction
}

func NewTransactionsHandler(service services.Transaction) *Transaction {
	return &Transaction{service: service}
}

func (c *Transaction) Create(ctx *gin.Context) {
	data, status, err := c.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *Transaction) Read(ctx *gin.Context) {
	data, status, err := c.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *Transaction) ReadById(ctx *gin.Context) {
	data, status, err := c.service.ReadById(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *Transaction) Update(ctx *gin.Context) {
	data, status, err := c.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *Transaction) Delete(ctx *gin.Context) {
	status, err := c.service.Delete(ctx)
	utils.SendResponse(ctx, commonsmodels.NewResponseSuccess(), status, err)
}

func (c *Transaction) Pay(ctx *gin.Context) {
	status, err := c.service.Pay(ctx)
	utils.SendResponse(ctx, commonsmodels.NewResponseSuccess(), status, err)
}
