package handlers

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/services"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type CreditCardsController struct {
	service services.CreditCard
}

func NewCreditCardsController(service services.CreditCard) *CreditCardsController {
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
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, err)
}
