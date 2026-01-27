package api

import (
	"financialcontrol/internal/handlers"
	"financialcontrol/internal/repositories"
	"financialcontrol/internal/services"
	"financialcontrol/internal/store"
	sm "financialcontrol/internal/store/models"
	"financialcontrol/internal/store/pgstore"
	crc "financialcontrol/internal/v1/creditcards/controllers"
	crs "financialcontrol/internal/v1/creditcards/services"
	cmtc "financialcontrol/internal/v1/monthly_transations/controllers"
	cmtr "financialcontrol/internal/v1/monthly_transations/repositories"
	cmts "financialcontrol/internal/v1/monthly_transations/services"
	ctc "financialcontrol/internal/v1/transactions/controllers"
	cts "financialcontrol/internal/v1/transactions/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router                        *gin.Engine
	categoriesHandler             *handlers.Category
	creditCardController          *crc.CreditCardsController
	transactionsController        *ctc.TransactionsController
	monthlyTransactionsController *cmtc.MonthlyTransactionsController
}

func NewApi(
	router *gin.Engine,
	pool *pgxpool.Pool,
) Api {
	store := pgstore.New(pool)
	repository := repositories.NewRepository(store)

	return Api{
		Router:                        router,
		categoriesHandler:             createCategory(repository),
		creditCardController:          createCreditCard(repository),
		transactionsController:        createTransactions(repository),
		monthlyTransactionsController: createMonthlyTransactions(store),
	}
}

func createCategory(repository repositories.Category) *handlers.Category {
	service := services.NewCategoriesService(repository)
	return handlers.NewCategoriesHandler(service)
}

func createCreditCard(repository sm.CreditCardsRepository) *crc.CreditCardsController {
	service := crs.NewCreditCardsService(repository)
	return crc.NewCreditCardsController(service)
}

func createTransactions(repository sm.TransactionsRepository) *ctc.TransactionsController {
	service := cts.NewTransactionsService(repository)
	return ctc.NewTransactionsController(service)
}

func createMonthlyTransactions(store store.MonthlyTransactionsStore) *cmtc.MonthlyTransactionsController {
	repository := cmtr.NewMonthlyTransactionsRepository(store)
	service := cmts.NewMonthlyTransactionsService(repository)
	return cmtc.NewMonthlyTransactionsController(service)
}
