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
	ctc "financialcontrol/internal/v1/transactions/controllers"
	ctr "financialcontrol/internal/v1/transactions/repositories"
	cts "financialcontrol/internal/v1/transactions/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router                 *gin.Engine
	categoriesController   *cac.CategoriesController
	creditCardController   *crc.CreditCardsController
	transactionsController *ctc.TransactionsController
}

func NewApi(
	router *gin.Engine,
	pool *pgxpool.Pool,
) Api {
	store := pgstore.New(pool)

	return Api{
		Router:                 router,
		categoriesController:   createCategory(store),
		creditCardController:   createCreditCard(store),
		transactionsController: createTransactions(store),
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

func createTransactions(store *pgstore.Queries) *ctc.TransactionsController {
	categoriesRepository := car.NewCategoriesRepository(store)
	creditcardRepository := crr.NewCreditCardsRepository(store)
	transactionsRepository := ctr.NewTransactionsRepository(store)
	service := cts.NewTransactionsService(categoriesRepository, creditcardRepository, transactionsRepository)
	return ctc.NewTransactionsController(service)
}
