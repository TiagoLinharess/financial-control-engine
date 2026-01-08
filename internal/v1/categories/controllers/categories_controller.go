package controllers

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	service categoriesModels.CategoriesService
}

func NewCategoriesController(service categoriesModels.CategoriesService) *CategoriesController {
	return &CategoriesController{service: service}
}

func (c *CategoriesController) Create(ctx *gin.Context) {
	data, status, err := c.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) Read(ctx *gin.Context) {
	data, status, err := c.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) ReadByID(ctx *gin.Context) {
	data, status, err := c.service.ReadByID(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) Update(ctx *gin.Context) {
	data, status, err := c.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) Delete(ctx *gin.Context) {
	status, err := c.service.Delete(ctx)
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, err)
}
