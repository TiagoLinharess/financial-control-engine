package controllers

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	categoriesModels "financialcontrol/internal/v1/categories/models"
	"net/http"
)

type CategoriesController struct {
	service categoriesModels.CategoriesService
}

func NewCategoriesController(service categoriesModels.CategoriesService) *CategoriesController {
	return &CategoriesController{service: service}
}

func (c *CategoriesController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.CreateCategory(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CategoriesController) ReadCategoriesByUser(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.ReadCategoriesByUser(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CategoriesController) ReadCategory(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.ReadCategoryByID(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CategoriesController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.UpdateCategory(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CategoriesController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	status, err := c.service.DeleteCategory(w, r)
	utils.SendResponse(w, models.NewResponseSuccess(), status, err)
}
