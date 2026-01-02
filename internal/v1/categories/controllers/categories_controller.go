package controllers

import (
	"financialcontrol/internal/utils"
	"financialcontrol/internal/v1/categories/models"
	"net/http"
)

type CategoriesController struct {
	service models.CategoriesService
}

func NewCategoriesController(service models.CategoriesService) *CategoriesController {
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
	utils.SendResponse(w, "success", http.StatusOK, nil)
}

func (c *CategoriesController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusOK, nil)
}
