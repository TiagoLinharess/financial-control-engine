package api

import (
	"financialcontrol/internal/store"
	"financialcontrol/internal/store/pgstore"
	cac "financialcontrol/internal/v1/categories/controllers"
	car "financialcontrol/internal/v1/categories/repositories"
	cas "financialcontrol/internal/v1/categories/services"
	crc "financialcontrol/internal/v1/creditcards/controllers"
	crr "financialcontrol/internal/v1/creditcards/repositories"
	crs "financialcontrol/internal/v1/creditcards/services"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router               *chi.Mux
	CategoriesController *cac.CategoriesController
	CreditCardController *crc.CreditCardsController
}

func NewApi(
	router *chi.Mux,
	pool *pgxpool.Pool,
) Api {
	store := pgstore.New(pool)

	return Api{
		Router:               router,
		CategoriesController: createCategory(store),
		CreditCardController: createCreditCard(store),
	}
}

func createCategory(store store.CategoriesStore) *cac.CategoriesController {
	repository := car.NewCategoriesRepository(store)
	service := cas.NewCategoriesService(repository)
	return cac.NewCategoriesController(service)
}

func createCreditCard(store store.CreditCardsStore) *crc.CreditCardsController {
	repository := crr.NewCreditCardsRepository(store)
	service := crs.NewCreditCardsService(repository)
	return crc.NewCreditCardsController(service)
}
