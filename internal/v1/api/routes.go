package api

func (a *Api) RegisterRoutes() {
	api := a.Router.Group("/engine/v1")

	categories := api.Group("/categories")
	{
		categories.POST("/", a.categoriesController.CreateCategory)
		categories.GET("/", a.categoriesController.ReadCategoriesByUser)
		categories.GET("/:id", a.categoriesController.ReadCategory)
		categories.PUT("/:id", a.categoriesController.UpdateCategory)
		categories.DELETE("/:id", a.categoriesController.DeleteCategory)
	}

	creditcards := api.Group("/creditcards")
	{
		creditcards.POST("/", a.creditCardController.Create)
		creditcards.GET("/", a.creditCardController.Read)
		creditcards.GET("/:id", a.creditCardController.ReadAt)
		creditcards.PUT("/:id", a.creditCardController.Update)
		creditcards.DELETE("/:id", a.creditCardController.Delete)
	}
}
