package api

import (
	"financialcontrol/internal/store"
	"financialcontrol/internal/store/pgstore"
	"financialcontrol/internal/v1/categories/controllers"
	"financialcontrol/internal/v1/categories/repositories"
	"financialcontrol/internal/v1/categories/services"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router               *chi.Mux
	CategoriesController *controllers.CategoriesController
}

func NewApi(
	router *chi.Mux,
	pool *pgxpool.Pool,
) Api {
	return Api{
		Router:               router,
		CategoriesController: createCategory(pgstore.New(pool)),
	}
}

func createCategory(store store.CategoriesStore) *controllers.CategoriesController {
	categoriesRepository := repositories.NewCategoriesRepository(store)
	categoriesService := services.NewCategoriesService(categoriesRepository)
	return controllers.NewCategoriesController(categoriesService)
}
