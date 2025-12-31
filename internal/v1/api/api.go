package api

import (
	"financialcontrol/internal/v1/categories/controllers"
	"financialcontrol/internal/v1/categories/repositories"
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
	return Api{
		Router:               router,
		CategoriesController: createCategory(),
	}
}

func createCategory() controllers.CategoriesController {
	categoriesRepository := repositories.NewCategoriesRepository()
	categoriesService := services.NewCategoriesService(categoriesRepository)
	return controllers.NewCategoriesController(categoriesService)
}
