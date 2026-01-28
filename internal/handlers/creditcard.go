package handlers

import (
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/services"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type CreditCard struct {
	service services.CreditCard
}

func NewCreditCardsHandler(service services.CreditCard) *CreditCard {
	return &CreditCard{service: service}
}

func (c *CreditCard) Create(ctx *gin.Context) {
	data, status, err := c.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCard) Read(ctx *gin.Context) {
	data, status, err := c.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCard) ReadAt(ctx *gin.Context) {
	data, status, err := c.service.ReadAt(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCard) Update(ctx *gin.Context) {
	data, status, err := c.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCard) Delete(ctx *gin.Context) {
	status, err := c.service.Delete(ctx)
	utils.SendResponse(ctx, commonsmodels.NewResponseSuccess(), status, err)
}
