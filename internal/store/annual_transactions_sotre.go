package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type AnnualTransactionsStore interface {
	CreateAnnualTransaction(ctx context.Context, arg pgstore.CreateAnnualTransactionParams) (pgstore.AnnualTransaction, error)
	ListAnnualTransactionsByUserIDPaginated(ctx context.Context, arg pgstore.ListAnnualTransactionsByUserIDPaginatedParams) ([]pgstore.ListAnnualTransactionsByUserIDPaginatedRow, error)
	GetAnnualTransactionByID(ctx context.Context, id uuid.UUID) (pgstore.AnnualTransaction, error)
	UpdateAnnualTransaction(ctx context.Context, arg pgstore.UpdateAnnualTransactionParams) (pgstore.AnnualTransaction, error)
	DeleteAnnualTransaction(ctx context.Context, id uuid.UUID) error
}
