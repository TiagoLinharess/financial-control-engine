package api

import "github.com/go-chi/chi/v5"

func (a *Api) RegisterRoutes() {
	a.Router.Route("/engine/v1", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Post("/", a.CategoriesController.CreateCategory)
			r.Get("/", a.CategoriesController.ReadCategories)
			r.Get("/{id}", a.CategoriesController.ReadCategory)
			r.Put("/{id}", a.CategoriesController.UpdateCategory)
			r.Delete("/{id}", a.CategoriesController.DeleteCategory)
		})
	})
}
