package api

func (a *Api) RegisterRoutes() {
	api := a.Router.Group("/engine/v1")

	categories := api.Group("/categories")
	{
		categories.POST("/", a.categoriesController.Create)
		categories.GET("/", a.categoriesController.Read)
		categories.GET("/:id", a.categoriesController.ReadByID)
		categories.PUT("/:id", a.categoriesController.Update)
		categories.DELETE("/:id", a.categoriesController.Delete)
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
	}
}
