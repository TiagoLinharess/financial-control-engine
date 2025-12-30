package api

import (
	"financialcontrol/internal/v1/categories/controllers"
	"financialcontrol/internal/v1/categories/services"

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router               *chi.Mux
	CategoriesController controllers.CategoriesController
}

func NewApi(
	router *chi.Mux,
) Api {
	categoriesService := services.NewCategoriesService()
	categoriesController := controllers.NewCategoriesController(categoriesService)

	return Api{
		Router:               router,
		CategoriesController: categoriesController,
	}
}
