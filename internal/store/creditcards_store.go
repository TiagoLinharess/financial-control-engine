package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type CreditCardsStore interface {
	CreateCreditCard(ctx context.Context, arg pgstore.CreateCreditCardParams) (pgstore.CreditCard, error)
	ListCreditCards(ctx context.Context, userID uuid.UUID) ([]pgstore.CreditCard, error)
	GetCreditCardByID(ctx context.Context, id uuid.UUID) (pgstore.CreditCard, error)
	CountCreditCardsByUserID(ctx context.Context, userID uuid.UUID) (int64, error)
	UpdateCreditCard(ctx context.Context, arg pgstore.UpdateCreditCardParams) (pgstore.CreditCard, error)
	DeleteCreditCard(ctx context.Context, id uuid.UUID) error
}
