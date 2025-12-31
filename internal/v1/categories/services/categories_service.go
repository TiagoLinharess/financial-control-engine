package services

import "financialcontrol/internal/v1/categories/models"

type CategoriesService struct {
	repository models.CategoriesRepository
}

func NewCategoriesService(repository models.CategoriesRepository) models.CategoriesService {
	return CategoriesService{repository: repository}
}
