package api

import (
	"financialcontrol/internal/handlers"
	"financialcontrol/internal/repositories"
	"financialcontrol/internal/services"
	"financialcontrol/internal/store/pgstore"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Api struct {
	Router                     *gin.Engine
	categoriesHandler          *handlers.Category
	creditCardsHandler         *handlers.CreditCard
	transactionsHandler        *handlers.Transaction
	monthlyTransactionsHandler *handlers.MonthlyTransaction
}

func NewApi(
	router *gin.Engine,
	pool *pgxpool.Pool,
) Api {
	store := pgstore.New(pool)
	repository := repositories.NewRepository(store)

	return Api{
		Router:                     router,
		categoriesHandler:          createCategory(repository),
		creditCardsHandler:         createCreditCard(repository),
		transactionsHandler:        createTransactions(repository),
		monthlyTransactionsHandler: createMonthlyTransactions(repository),
	}
}

func createCategory(repository repositories.Category) *handlers.Category {
	service := services.NewCategoriesService(repository)
	return handlers.NewCategoriesHandler(service)
}

func createCreditCard(repository repositories.CreditCard) *handlers.CreditCard {
	service := services.NewCreditCardsService(repository)
	return handlers.NewCreditCardsHandler(service)
}

func createTransactions(repository repositories.Transaction) *handlers.Transaction {
	service := services.NewTransactionsService(repository)
	return handlers.NewTransactionsHandler(service)
}

func createMonthlyTransactions(repository repositories.MonthlyTransaction) *handlers.MonthlyTransaction {
	service := services.NewMonthlyTransactionService(repository)
	return handlers.NewMonthlyTransactionsHandler(service)
}
