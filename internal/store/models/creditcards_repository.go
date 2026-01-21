package models

import (
	"context"
	e "financialcontrol/internal/models/errors"
	cr "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

type CreditCardsRepository interface {
	CreateCreditCard(context context.Context, creditCard cr.CreateCreditCard) (cr.CreditCard, []e.ApiError)
	ReadCreditCards(context context.Context, userId uuid.UUID) ([]cr.CreditCard, []e.ApiError)
	ReadCountByUser(context context.Context, userId uuid.UUID) (int, []e.ApiError)
	ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (cr.CreditCard, []e.ApiError)
	UpdateCreditCard(context context.Context, creditCard cr.CreditCard) (cr.CreditCard, []e.ApiError)
	DeleteCreditCard(context context.Context, creditCardId uuid.UUID) []e.ApiError
	HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, []e.ApiError)
}
