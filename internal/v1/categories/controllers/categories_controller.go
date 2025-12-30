package controllers

import (
	"financialcontrol/internal/utils"
	"financialcontrol/internal/v1/categories/models"
	"net/http"
)

type CategoriesController struct {
	service models.CategoriesService
}

func NewCategoriesController(service models.CategoriesService) CategoriesController {
	return CategoriesController{service: service}
}

func (c *CategoriesController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	categoryResponse, err := c.service.CreateCategory(w, r)

	if len(err.Errors) > 0 {
		utils.SendError(w, err)
		return
	}

	utils.SendResponse(w, categoryResponse, http.StatusCreated)
}

func (c *CategoriesController) ReadCategories(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusOK)
}

func (c *CategoriesController) ReadCategory(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusOK)
}

func (c *CategoriesController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusOK)
}

func (c *CategoriesController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusOK)
}
