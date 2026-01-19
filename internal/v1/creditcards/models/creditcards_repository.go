package models

import (
	"context"
	"financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type CreditCardsRepository interface {
	Create(context context.Context, creditCard CreateCreditCard) (CreditCard, []errors.ApiError)
	Read(context context.Context, userId uuid.UUID) ([]CreditCard, []errors.ApiError)
	ReadCountByUser(context context.Context, userId uuid.UUID) (int, []errors.ApiError)
	ReadByID(context context.Context, creditCardId uuid.UUID) (CreditCard, []errors.ApiError)
	Update(context context.Context, creditCard CreditCard) (CreditCard, []errors.ApiError)
	Delete(context context.Context, creditCardId uuid.UUID) []errors.ApiError
	HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, []errors.ApiError)
}
