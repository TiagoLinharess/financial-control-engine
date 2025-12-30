package api

import (
	"financialcontrol/internal/v1/categories/controllers"

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
		CategoriesController: controllers.NewCategoriesController(),
	}
}
