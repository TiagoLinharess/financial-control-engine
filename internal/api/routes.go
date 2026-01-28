package api

import "financialcontrol/internal/middlewares"

func (a *Api) RegisterRoutes() {
	api := a.Router.Group("/engine/v1")

	api.Use(middlewares.UserIDMiddleware())

	categories := api.Group("/categories")
	{
		categories.POST("/", a.categoriesHandler.Create())
		categories.GET("/", a.categoriesHandler.Read())
		categories.GET("/:id", a.categoriesHandler.ReadByID())
		categories.PUT("/:id", a.categoriesHandler.Update())
		categories.DELETE("/:id", a.categoriesHandler.Delete())
	}

	creditcards := api.Group("/creditcards")
	{
		creditcards.POST("/", a.creditCardsHandler.Create)
		creditcards.GET("/", a.creditCardsHandler.Read)
		creditcards.GET("/:id", a.creditCardsHandler.ReadAt)
		creditcards.PUT("/:id", a.creditCardsHandler.Update)
		creditcards.DELETE("/:id", a.creditCardsHandler.Delete)
	}

	trasactions := api.Group("/transactions")
	{
		trasactions.POST("/", a.transactionsHandler.Create)
		trasactions.GET("/", a.transactionsHandler.Read)
		trasactions.GET("/:id", a.transactionsHandler.ReadById)
		trasactions.PUT("/:id", a.transactionsHandler.Update)
		trasactions.DELETE("/:id", a.transactionsHandler.Delete)
		trasactions.PUT("/pay/:id", a.transactionsHandler.Pay)
	}

	monthlyTransactions := api.Group("/monthly_transactions")
	{
		monthlyTransactions.POST("/", a.monthlyTransactionsHandler.Create)
		monthlyTransactions.GET("/", a.monthlyTransactionsHandler.Read)
		monthlyTransactions.GET("/:id", a.monthlyTransactionsHandler.ReadById)
		monthlyTransactions.PUT("/:id", a.monthlyTransactionsHandler.Update)
		monthlyTransactions.DELETE("/:id", a.monthlyTransactionsHandler.Delete)
	}
}
