package services

import "financialcontrol/internal/v1/categories/models"

type CategoriesService struct {
}

func NewCategoriesService() models.CategoriesService {
	return CategoriesService{}
}
