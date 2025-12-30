package controllers

import (
	"financialcontrol/internal/utils"
	"net/http"
)

type CategoriesController struct {
}

func NewCategoriesController() CategoriesController {
	return CategoriesController{}
}

func (c *CategoriesController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	utils.SendResponse(w, "success", http.StatusCreated)
}
