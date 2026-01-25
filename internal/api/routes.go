package api

func (a *Api) RegisterRoutes() {
	api := a.Router.Group("/engine/v1")

	categories := api.Group("/categories")
	{
		categories.POST("/", a.categoriesHandler.Create)
		categories.GET("/", a.categoriesHandler.Read)
		categories.GET("/:id", a.categoriesHandler.ReadByID)
		categories.PUT("/:id", a.categoriesHandler.Update)
		categories.DELETE("/:id", a.categoriesHandler.Delete)
	}

	creditcards := api.Group("/creditcards")
	{
		creditcards.POST("/", a.creditCardController.Create)
		creditcards.GET("/", a.creditCardController.Read)
		creditcards.GET("/:id", a.creditCardController.ReadAt)
		creditcards.PUT("/:id", a.creditCardController.Update)
		creditcards.DELETE("/:id", a.creditCardController.Delete)
	}

	trasactions := api.Group("/transactions")
	{
		trasactions.POST("/", a.transactionsController.Create)
		trasactions.GET("/", a.transactionsController.Read)
		trasactions.GET("/:id", a.transactionsController.ReadById)
		trasactions.PUT("/:id", a.transactionsController.Update)
		trasactions.DELETE("/:id", a.transactionsController.Delete)
		trasactions.PUT("/pay/:id", a.transactionsController.Pay)
	}

	monthlyTransactions := api.Group("/monthly_transactions")
	{
		monthlyTransactions.POST("/", a.monthlyTransactionsController.Create)
		monthlyTransactions.GET("/", a.monthlyTransactionsController.Read)
		monthlyTransactions.GET("/:id", a.monthlyTransactionsController.ReadById)
		monthlyTransactions.PUT("/:id", a.monthlyTransactionsController.Update)
		monthlyTransactions.DELETE("/:id", a.monthlyTransactionsController.Delete)
	}
}
