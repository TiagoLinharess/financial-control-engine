package controllers

import (
	m "financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"

	"github.com/gin-gonic/gin"
)

type CreditCardsController struct {
	service cm.CreditCardsService
}

func NewCreditCardsController(service cm.CreditCardsService) *CreditCardsController {
	return &CreditCardsController{service: service}
}

func (c *CreditCardsController) Create(ctx *gin.Context) {
	data, status, err := c.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCardsController) Read(ctx *gin.Context) {
	data, status, err := c.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCardsController) ReadAt(ctx *gin.Context) {
	data, status, err := c.service.ReadAt(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCardsController) Update(ctx *gin.Context) {
	data, status, err := c.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CreditCardsController) Delete(ctx *gin.Context) {
	status, err := c.service.Delete(ctx)
	utils.SendResponse(ctx, m.NewResponseSuccess(), status, err)
}
