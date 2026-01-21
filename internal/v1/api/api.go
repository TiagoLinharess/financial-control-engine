package api

import (
	"financialcontrol/internal/store"
	sm "financialcontrol/internal/store/models"
	"financialcontrol/internal/store/pgstore"
	r "financialcontrol/internal/store/repositories"
	cac "financialcontrol/internal/v1/categories/controllers"
	cas "financialcontrol/internal/v1/categories/services"
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
	categoriesController          *cac.CategoriesController
	creditCardController          *crc.CreditCardsController
	transactionsController        *ctc.TransactionsController
	monthlyTransactionsController *cmtc.MonthlyTransactionsController
}

func NewApi(
	router *gin.Engine,
	pool *pgxpool.Pool,
) Api {
	store := pgstore.New(pool)
	repository := r.NewRepository(store)

	return Api{
		Router:                        router,
		categoriesController:          createCategory(repository),
		creditCardController:          createCreditCard(repository),
		transactionsController:        createTransactions(repository),
		monthlyTransactionsController: createMonthlyTransactions(store),
	}
}

func createCategory(repository sm.CategoriesRepository) *cac.CategoriesController {
	service := cas.NewCategoriesService(repository)
	return cac.NewCategoriesController(service)
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
