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

func (c *CategoriesController) CreateCategory(ctx *gin.Context) {
	data, status, err := c.service.CreateCategory(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) ReadCategoriesByUser(ctx *gin.Context) {
	data, status, err := c.service.ReadCategoriesByUser(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) ReadCategory(ctx *gin.Context) {
	data, status, err := c.service.ReadCategoryByID(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) UpdateCategory(ctx *gin.Context) {
	data, status, err := c.service.UpdateCategory(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (c *CategoriesController) DeleteCategory(ctx *gin.Context) {
	status, err := c.service.DeleteCategory(ctx)
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, err)
}
