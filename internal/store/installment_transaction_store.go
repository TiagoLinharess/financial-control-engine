package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type InstallmentTransactionsStore interface {
	CreateInstallmentTransaction(ctx context.Context, arg pgstore.CreateInstallmentTransactionParams) (pgstore.InstallmentTransaction, error)
	ListInstallmentTransactionsByUserIDPaginated(ctx context.Context, arg pgstore.ListInstallmentTransactionsByUserIDPaginatedParams) ([]pgstore.ListInstallmentTransactionsByUserIDPaginatedRow, error)
	GetInstallmentTransactionByID(ctx context.Context, id uuid.UUID) (pgstore.InstallmentTransaction, error)
	UpdateInstallmentTransaction(ctx context.Context, arg pgstore.UpdateInstallmentTransactionParams) (pgstore.InstallmentTransaction, error)
	DeleteInstallmentTransaction(ctx context.Context, id uuid.UUID) error
}
