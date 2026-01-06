package services

import "financialcontrol/internal/v1/creditcards/models"

type CreditCardsService struct {
	repository models.CreditCardsRepository
}

func NewCreditCardsService(repository models.CreditCardsRepository) models.CreditCardsService {
	return &CreditCardsService{repository: repository}
}
